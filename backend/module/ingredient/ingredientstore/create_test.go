package ingredientstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateIngredient(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *ingredientmodel.IngredientCreate
	}

	id := "ingredient1"
	name := "Ingredient 1"
	measureType := enum.Weight
	price := float32(1000)
	ingredientCreate := ingredientmodel.IngredientCreate{
		Id:          &id,
		Name:        name,
		MeasureType: &measureType,
		Price:       price,
	}
	mockErr := errors.New(mock.Anything)
	mockErrName := &common.GormErr{
		Number:  1062,
		Message: "name",
	}
	mockErrPRIMARY := &common.GormErr{
		Number:  1062,
		Message: "PRIMARY",
	}
	expectedSql := "INSERT INTO `Ingredient` (`id`,`name`,`measureType`,`price`) VALUES (?,?,?,?)"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    error
		wantErr bool
	}{
		{
			name: "Create ingredient in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredientCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(id, name, measureType, price).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create ingredient in database failed because of database error",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredientCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(id, name, measureType, price).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Create ingredient in database failed because of duplicate name",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredientCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(id, name, measureType, price).
					WillReturnError(mockErrName)
				sqlDBMock.ExpectRollback()
			},
			want:    ingredientmodel.ErrIngredientNameDuplicate,
			wantErr: true,
		},
		{
			name: "Create ingredient in database failed because of duplicate id",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &ingredientCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(id, name, measureType, price).
					WillReturnError(mockErrPRIMARY)
				sqlDBMock.ExpectRollback()
			},
			want:    ingredientmodel.ErrIngredientIdDuplicate,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateIngredient(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateIngredient() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateIngredient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
