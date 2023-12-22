package supplierstore

import (
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_GetAllSupplier(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}

	mockData := []suppliermodel.SimpleSupplier{
		{
			Id:    "Supplier001",
			Name:  "SupplierName001",
			Phone: "0123456789",
		},
		{
			Id:    "Supplier002",
			Name:  "SupplierName002",
			Phone: "0123456789",
		},
	}

	expectedQuery := "SELECT * FROM `Supplier` ORDER BY name"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		mock    func()
		want    []suppliermodel.SimpleSupplier
		wantErr bool
	}{
		{
			name: "Get all suppliers successfully",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "phone"})

				for _, data := range mockData {
					rows.AddRow(data.Id, data.Name, data.Phone)
				}

				sqlDBMock.ExpectQuery(expectedQuery).WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Get all suppliers failed",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).WillReturnError(mockErr)
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

			got, err := s.GetAllSupplier(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got, "GetAllSupplier() got = %v, want %v", got, tt.want)
		})
	}
}
