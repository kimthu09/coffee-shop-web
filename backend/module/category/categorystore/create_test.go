package categorystore

import (
	"coffee_shop_management_backend/common"
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

type FakeGormErr struct {
	Number  int    `json:"Numbers"`
	Message string `json:"Messages"`
}

func (gErr *FakeGormErr) Error() string {
	return gErr.Message
}

func Test_sqlStore_CreateCategory(t *testing.T) {
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

	categoryId := mock.Anything
	categoryName := mock.Anything
	categoryDescription := mock.Anything
	categoryCreate := categorymodel.CategoryCreate{
		Id:          categoryId,
		Name:        categoryName,
		Description: categoryDescription,
	}
	mockErr := errors.New("some thing when wrong")
	mockErrName := &common.GormErr{
		Number:  1062,
		Message: "name",
	}
	mockErrPRIMARY := &common.GormErr{
		Number:  1062,
		Message: "PRIMARY",
	}
	mockErrFaKeGorm := &FakeGormErr{
		Number:  1062,
		Message: "FakeGorm",
	}
	expectedSql := "INSERT INTO `Category` (`id`,`name`,`description`) VALUES (?,?,?)"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data *categorymodel.CategoryCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    error
		wantErr bool
	}{
		{
			name: "Create category in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						categoryCreate.Id,
						categoryCreate.Name,
						categoryCreate.Description).
					WillReturnResult(sqlmock.NewResult(1, 1))
				sqlDBMock.ExpectCommit()
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Create category in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						categoryCreate.Id,
						categoryCreate.Name,
						categoryCreate.Description).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErr),
			wantErr: true,
		},
		{
			name: "Create category in database failed because duplicate id",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						categoryCreate.Id,
						categoryCreate.Name,
						categoryCreate.Description).
					WillReturnError(mockErrPRIMARY)
				sqlDBMock.ExpectRollback()
			},
			want:    categorymodel.ErrCategoryIdDuplicate,
			wantErr: true,
		},
		{
			name: "Create category in database failed because duplicate name",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						categoryCreate.Id,
						categoryCreate.Name,
						categoryCreate.Description).
					WillReturnError(mockErrName)
				sqlDBMock.ExpectRollback()
			},
			want:    categorymodel.ErrCategoryNameDuplicate,
			wantErr: true,
		},
		{
			name: "Create category in database failed because error is not in right format",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: &categoryCreate,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						categoryCreate.Id,
						categoryCreate.Name,
						categoryCreate.Description).
					WillReturnError(mockErrFaKeGorm)
				sqlDBMock.ExpectRollback()
			},
			want:    common.ErrDB(mockErrFaKeGorm),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateCategory(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, err, tt.want, "CreateCategory() = %v, want %v", err, tt.want)
			} else {
				assert.Nil(t, err, "CreateCategory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
