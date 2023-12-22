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

func Test_sqlStore_UpdateInfoCategory(t *testing.T) {
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
	nameUpdate := mock.Anything
	descriptionUpdate := mock.Anything
	categoryUpdateInfo := categorymodel.CategoryUpdateInfo{
		Name:        &nameUpdate,
		Description: &descriptionUpdate,
	}

	expectedSql := "UPDATE `Category` SET `name`=?,`description`=? WHERE id = ?"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		id   string
		data *categorymodel.CategoryUpdateInfo
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Update category info successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdateInfo,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(categoryUpdateInfo.Name, categoryUpdateInfo.Description, id).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Update category info failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				id:   id,
				data: &categoryUpdateInfo,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedSql).
					WithArgs(categoryUpdateInfo.Name, categoryUpdateInfo.Description, id).
					WillReturnError(errors.New("some error"))
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

			err := s.UpdateInfoCategory(tt.args.ctx, tt.args.id, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "UpdateInfoCategory() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "UpdateInfoCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
