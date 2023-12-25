package shopgeneralstore

import (
	"coffee_shop_management_backend/module/shopgeneral/shopgeneralmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindShopGeneral(t *testing.T) {
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

	// Mock shop general data
	shopGeneralName := "MockShopGeneral"
	shopGeneralEmail := "mock@example.com"
	shopGeneralPhone := "123456789"
	shopGeneralAddress := "Mock Address"
	shopGeneralWifiPass := "MockWifiPass"
	shopGeneralAccumulatePointPercent := float32(0.1)
	shopGeneralUsePointPercent := float32(0.2)

	shopGeneral := shopgeneralmodel.ShopGeneral{
		Name:                   shopGeneralName,
		Email:                  shopGeneralEmail,
		Phone:                  shopGeneralPhone,
		Address:                shopGeneralAddress,
		WifiPass:               shopGeneralWifiPass,
		AccumulatePointPercent: shopGeneralAccumulatePointPercent,
		UsePointPercent:        shopGeneralUsePointPercent,
	}

	mockErr := errors.New("something went wrong with the database")

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}

	expectedSql := "SELECT * FROM `ShopGeneral` ORDER BY `ShopGeneral`.`name` LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *shopgeneralmodel.ShopGeneral
		wantErr bool
	}{
		{
			name: "Find shop general in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WillReturnRows(
						sqlmock.NewRows([]string{"name", "email", "phone", "address", "wifiPass", "accumulatePointPercent", "usePointPercent"}).
							AddRow(
								shopGeneral.Name,
								shopGeneral.Email,
								shopGeneral.Phone,
								shopGeneral.Address,
								shopGeneral.WifiPass,
								shopGeneral.AccumulatePointPercent,
								shopGeneral.UsePointPercent,
							),
					)
			},
			want:    &shopGeneral,
			wantErr: false,
		},
		{
			name: "Find shop general in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find shop general in database record not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindShopGeneral(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(t, err, "FindShopGeneral() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindShopGeneral() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindShopGeneral() got = %v, want %v", got, tt.want)
			}
		})
	}
}
