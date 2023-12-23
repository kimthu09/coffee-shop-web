package userstore

import (
	"coffee_shop_management_backend/module/user/usermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_GetAllUser(t *testing.T) {
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

	mockData := []usermodel.SimpleUser{
		{
			Id:    "User001",
			Name:  "UserName001",
			Email: "user001@gmail.com",
		},
		{
			Id:    "User002",
			Name:  "UserName002",
			Email: "user002@gmail.com",
		},
	}

	expectedQuery := "SELECT * FROM `MUser` ORDER BY name"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		mock    func()
		want    []usermodel.SimpleUser
		wantErr bool
	}{
		{
			name: "Get all users successfully",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email"})

				for _, data := range mockData {
					rows.AddRow(data.Id, data.Name, data.Email)
				}

				sqlDBMock.ExpectQuery(expectedQuery).WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Get all users failed",
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

			got, err := s.GetAllUser(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "GetAllSupplier() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "GetAllSupplier() got = %v, want %v", got, tt.want)
			}
		})
	}
}
