package shopgeneralstore

import (
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_UpdateGeneralShop(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	name := "UpdatedShop"
	email := "updated@example.com"
	phone := "0987654321"
	address := "Updated Address"
	wifiPass := "Updated WifiPass"
	accumulatePointPercent := float32(0.002)
	usePointPercentPercent := float32(1)
	updatedShopGeneral := &shopgeneralmodel.ShopGeneralUpdate{
		Name:                   &name,
		Email:                  &email,
		Phone:                  &phone,
		Address:                &address,
		WifiPass:               &wifiPass,
		AccumulatePointPercent: &accumulatePointPercent,
		UsePointPercent:        &usePointPercentPercent,
	}

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *shopgeneralmodel.ShopGeneralUpdate
	}

	expectedSql := "UPDATE `ShopGeneral` SET `name`=?,`email`=?,`phone`=?,`address`=?,`wifiPass`=?,`accumulatePointPercent`=?,`usePointPercent`=? WHERE id = ?"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update shop general in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: updatedShopGeneral,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						updatedShopGeneral.Name,
						updatedShopGeneral.Email,
						updatedShopGeneral.Phone,
						updatedShopGeneral.Address,
						updatedShopGeneral.WifiPass,
						updatedShopGeneral.AccumulatePointPercent,
						updatedShopGeneral.UsePointPercent,
						"shop",
					).
					WillReturnResult(sqlmock.NewResult(0, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update shop general in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: updatedShopGeneral,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						updatedShopGeneral.Name,
						updatedShopGeneral.Email,
						updatedShopGeneral.Phone,
						updatedShopGeneral.Address,
						updatedShopGeneral.WifiPass,
						updatedShopGeneral.AccumulatePointPercent,
						updatedShopGeneral.UsePointPercent,
						"shop",
					).
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

			err := s.UpdateGeneralShop(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateGeneralShop() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateGeneralShop() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
