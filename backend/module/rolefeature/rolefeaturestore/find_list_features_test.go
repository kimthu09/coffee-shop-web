package rolefeaturestore

import (
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

func Test_sqlStore_FindListFeatures(t *testing.T) {
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
	type args struct {
		ctx    context.Context
		roleId string
	}

	roleId := mock.Anything
	feature1 := mock.Anything
	feature2 := mock.Anything
	expectedSql := "SELECT * FROM `RoleFeature` WHERE roleId = ?"
	roleFeatures := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    roleId,
			FeatureId: feature1,
		},
		{
			RoleId:    roleId,
			FeatureId: feature2,
		},
	}
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []rolefeaturemodel.RoleFeature
		wantErr bool
	}{
		{
			name: "Find list feature by role successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						roleId).
					WillReturnRows(
						sqlmock.
							NewRows(
								[]string{"roleId", "featureId"}).
							AddRow(roleId, feature1).
							AddRow(roleId, feature2),
					)
			},
			want:    roleFeatures,
			wantErr: false,
		},
		{
			name: "Find list feature by role failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:    context.Background(),
				roleId: roleId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedSql).
					WithArgs(
						roleId).
					WillReturnError(mockErr)
			},
			want:    roleFeatures,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindListFeatures(tt.args.ctx, tt.args.roleId)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListFeatures() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListFeatures() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListFeatures() got = %v, want %v", got, tt.want)
			}
		})
	}
}
