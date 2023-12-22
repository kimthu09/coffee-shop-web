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

func Test_sqlStore_CreateListRoleFeatureDetail(t *testing.T) {
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
		ctx  context.Context
		data []rolefeaturemodel.RoleFeature
	}

	mockRoleFeatures := []rolefeaturemodel.RoleFeature{
		{
			RoleId:    "Role001",
			FeatureId: "Feature001",
		},
		{
			RoleId:    "Role002",
			FeatureId: "Feature002",
		},
	}
	mockErr := errors.New(mock.Anything)
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create role feature in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockRoleFeatures,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec("INSERT INTO `RoleFeature` (`roleId`,`featureId`) VALUES (?,?),(?,?)").
					WithArgs(mockRoleFeatures[0].RoleId, mockRoleFeatures[0].FeatureId, mockRoleFeatures[1].RoleId, mockRoleFeatures[1].FeatureId).
					WillReturnResult(sqlmock.NewResult(1, int64(len(mockRoleFeatures))))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create role feature in database failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockRoleFeatures,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec("INSERT INTO `RoleFeature` (`roleId`,`featureId`) VALUES (?,?),(?,?)").
					WithArgs(mockRoleFeatures[0].RoleId, mockRoleFeatures[0].FeatureId, mockRoleFeatures[1].RoleId, mockRoleFeatures[1].FeatureId).
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

			err := s.CreateListRoleFeatureDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListRoleFeatureDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListRoleFeatureDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
