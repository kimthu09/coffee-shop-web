package productstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/product/productmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateTopping(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	id := "product001"
	toppingCreate := &productmodel.ToppingCreate{
		ProductCreate: &productmodel.ProductCreate{
			Id:           &id,
			Name:         "TestTopping",
			Description:  "Description for TestTopping",
			CookingGuide: "Cooking guide for TestTopping",
		},
		Cost:     10000,
		Price:    8000,
		RecipeId: "Recipe001",
	}
	expectedQuery := "INSERT INTO `Topping` (`id`,`name`,`description`,`cookingGuide`,`cost`,`price`,`recipeId`) VALUES (?,?,?,?,?,?,?)"
	mockErr := errors.New(mock.Anything)
	mockErrName := &common.GormErr{
		Number:  1062,
		Message: "name",
	}
	mockErrPRIMARY := &common.GormErr{
		Number:  1062,
		Message: "PRIMARY",
	}

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *productmodel.ToppingCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    error
		wantErr bool
	}{
		{
			name:   "Create topping successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: toppingCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						toppingCreate.Id,
						toppingCreate.Name,
						toppingCreate.Description,
						toppingCreate.CookingGuide,
						toppingCreate.Cost,
						toppingCreate.Price,
						toppingCreate.RecipeId,
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mockSqlDB.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:   "Create topping failed because can not save data to database",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: toppingCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						toppingCreate.Id,
						toppingCreate.Name,
						toppingCreate.Description,
						toppingCreate.CookingGuide,
						toppingCreate.Cost,
						toppingCreate.Price,
						toppingCreate.RecipeId,
					).
					WillReturnError(mockErr)
				mockSqlDB.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name:   "Create topping failed because duplicate id",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: toppingCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						toppingCreate.Id,
						toppingCreate.Name,
						toppingCreate.Description,
						toppingCreate.CookingGuide,
						toppingCreate.Cost,
						toppingCreate.Price,
						toppingCreate.RecipeId,
					).
					WillReturnError(mockErrPRIMARY)
				mockSqlDB.ExpectRollback()
			},
			want:    productmodel.ErrToppingIdDuplicate,
			wantErr: true,
		},
		{
			name:   "Create topping failed because duplicate name",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: toppingCreate,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WithArgs(
						toppingCreate.Id,
						toppingCreate.Name,
						toppingCreate.Description,
						toppingCreate.CookingGuide,
						toppingCreate.Cost,
						toppingCreate.Price,
						toppingCreate.RecipeId,
					).
					WillReturnError(mockErrName)
				mockSqlDB.ExpectRollback()
			},
			want:    productmodel.ErrToppingNameDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateTopping(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateTopping() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateTopping() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateTopping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
