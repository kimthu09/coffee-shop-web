package rolestore

import (
	"coffee_shop_management_backend/module/role/rolemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindRole(t *testing.T) {
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

	roleId := mock.Anything
	mockConditions := map[string]interface{}{
		"id": roleId,
	}
	expectedQuery := "SELECT * FROM `Role` WHERE `id` = ? ORDER BY `Role`.`id` LIMIT 1"
	role := rolemodel.Role{
		Id:   roleId,
		Name: "Role001",
	}
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *rolemodel.Role
		wantErr bool
	}{
		{
			name: "Find role in database successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["id"]).
					WillReturnRows(
						sqlmock.
							NewRows([]string{"id", "name"}).
							AddRow(role.Id, role.Name),
					)
			},
			want:    &role,
			wantErr: false,
		},
		{
			name: "Find role in database failed something went wrong with the database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["id"]).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find role in database record not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedQuery).
					WithArgs(mockConditions["foodId"], mockConditions["sizeId"]).
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

			got, err := s.FindRole(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindRole() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindRole() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindRole() got = %v, want %v", got, tt.want)
			}
		})
	}
}
