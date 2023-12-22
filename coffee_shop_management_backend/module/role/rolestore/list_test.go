package rolestore

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListRole(t *testing.T) {
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

	mockData := []rolemodel.Role{
		{
			Id:   "Role001",
			Name: "Admin",
			RoleFeatures: []rolefeaturemodel.RoleFeature{
				{
					RoleId:    "Role001",
					FeatureId: "Feature001",
				},
				{
					RoleId:    "Role001",
					FeatureId: "Feature002",
				},
			},
		},
		{
			Id:   "Role002",
			Name: "User",
			RoleFeatures: []rolefeaturemodel.RoleFeature{
				{
					RoleId:    "Role002",
					FeatureId: "Feature003",
				},
				{
					RoleId:    "Role002",
					FeatureId: "Feature004",
				},
			},
		},
	}

	expectedRoleQuery := "SELECT * FROM `Role`"
	expectedRoleFeatureQuery := "SELECT * FROM `RoleFeature` WHERE `RoleFeature`.`roleId` IN (?,?)"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		mock    func()
		want    []rolemodel.Role
		wantErr bool
	}{
		{
			name: "List roles successfully",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(mockData[0].Id, mockData[0].Name).
					AddRow(mockData[1].Id, mockData[1].Name)

				sqlDBMock.ExpectQuery(expectedRoleQuery).
					WillReturnRows(rows)

				roleFeaturesRows := sqlmock.NewRows([]string{"roleId", "featureId"}).
					AddRow(
						mockData[0].RoleFeatures[0].RoleId,
						mockData[0].RoleFeatures[0].FeatureId,
					).
					AddRow(
						mockData[0].RoleFeatures[1].RoleId,
						mockData[0].RoleFeatures[1].FeatureId,
					).
					AddRow(
						mockData[1].RoleFeatures[0].RoleId,
						mockData[1].RoleFeatures[0].FeatureId,
					).
					AddRow(
						mockData[1].RoleFeatures[1].RoleId,
						mockData[1].RoleFeatures[1].FeatureId,
					)

				sqlDBMock.ExpectQuery(expectedRoleFeatureQuery).
					WithArgs(
						mockData[0].RoleFeatures[1].RoleId,
						mockData[1].RoleFeatures[1].RoleId,
					).
					WillReturnRows(roleFeaturesRows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "List roles failed",
			fields: fields{
				db: gormDB,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedRoleQuery).
					WillReturnError(mockErr)
			},
			want:    mockData,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListRole(context.Background())

			if tt.wantErr {
				assert.NotNil(t, err, "ListRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListRole() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}
