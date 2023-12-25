package sizefoodstore

import (
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateSizeFood(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err) // Error here
	}

	sizeFoodCreate := sizefoodmodel.SizeFoodCreate{
		FoodId:   mock.Anything,
		SizeId:   mock.Anything,
		Name:     mock.Anything,
		Cost:     0,
		Price:    0,
		RecipeId: mock.Anything,
	}
	expectedSql := "INSERT INTO `SizeFood` (`foodId`,`sizeId`,`name`,`cost`,`price`,`recipeId`) VALUES (?,?,?,?,?,?)"
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *sizefoodmodel.SizeFoodCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create size food in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &sizeFoodCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						sizeFoodCreate.FoodId,
						sizeFoodCreate.SizeId,
						sizeFoodCreate.Name,
						sizeFoodCreate.Cost,
						sizeFoodCreate.Price,
						sizeFoodCreate.RecipeId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create size food in database failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &sizeFoodCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						sizeFoodCreate.FoodId,
						sizeFoodCreate.SizeId,
						sizeFoodCreate.Name,
						sizeFoodCreate.Cost,
						sizeFoodCreate.Price,
						sizeFoodCreate.RecipeId).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateSizeFood(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateSizeFood() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
