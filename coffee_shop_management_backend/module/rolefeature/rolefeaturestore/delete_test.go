package rolefeaturestore

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_DeleteRoleFeature(t *testing.T) {
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
		ctx        context.Context
		conditions map[string]interface{}
	}

	mockConditions := map[string]interface{}{
		"roleId":    "Role001",
		"featureId": "Feature001",
	}
	expectedQuery := "DELETE FROM `RoleFeature` WHERE `featureId` = ? AND `roleId` = ?"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Delete role feature successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockConditions["featureId"], mockConditions["roleId"]).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Delete role feature failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(mockConditions["featureId"], mockConditions["roleId"]).
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

			err := s.DeleteRoleFeature(tt.args.ctx, tt.args.conditions)

			if tt.wantErr {
				assert.NotNil(t, err, "DeleteRoleFeature() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "DeleteRoleFeature() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
