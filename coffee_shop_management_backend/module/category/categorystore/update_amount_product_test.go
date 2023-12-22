package categorystore

import (
	"coffee_shop_management_backend/module/category/categorymodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateAmountProductCategory(t *testing.T) {
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

	id := "1"
	amountProduct := 10
	categoryUpdateAmountProduct := categorymodel.CategoryUpdateAmountProduct{
		AmountProduct: &amountProduct,
	}
	mockErr := errors.New(mock.Anything)
	expectedSql := "UPDATE `Category` SET `amountProduct`=amountProduct + ? WHERE id = ?"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *categorymodel.CategoryUpdateAmountProduct
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update amount of products in category successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdateAmountProduct,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(amountProduct, id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update amount of products in category failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdateAmountProduct,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(amountProduct, id).
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

			err := s.UpdateAmountProductCategory(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateAmountProductCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateAmountProductCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
