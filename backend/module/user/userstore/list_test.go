package userstore

import (
	"coffee_shop_management_backend/common"
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

func Test_sqlStore_ListUser(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	userSearch := "user123"
	active := true
	inactive := false
	filterActive := &usermodel.Filter{
		SearchKey: "john",
		IsActive:  &active,
		Role:      "admin",
	}
	filterInactive := &usermodel.Filter{
		SearchKey: "john",
		IsActive:  &inactive,
		Role:      "admin",
	}
	mockProperties := []string{"id", "name", "email", "phone", "address"}
	paging := &common.Paging{
		Page:  1,
		Limit: 10,
	}
	listUsers := []usermodel.User{
		{
			Id:       "1",
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "hashed_password",
			Salt:     "salt",
			RoleId:   "admin",
			IsActive: true,
		},
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"email",
		"password",
		"salt",
		"roleId",
		"isActive",
	})
	for _, user := range listUsers {
		rows.AddRow(user.Id, user.Name, user.Email, user.Password, user.Salt, user.RoleId, user.IsActive)
	}

	queryString :=
		"SELECT * FROM `MUser` " +
			"WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ? OR address LIKE ?) " +
			"AND isActive = ? " +
			"AND roleId = ? " +
			"AND id <> ? " +
			"ORDER BY name LIMIT 10"

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listUsers))
	queryStringCount :=
		"SELECT count(*) FROM `MUser` " +
			"WHERE (id LIKE ? OR name LIKE ? OR email LIKE ? OR phone LIKE ? OR address LIKE ?) " +
			"AND isActive = ? " +
			"AND roleId = ? " +
			"AND id <> ?"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		userSearch                 string
		filter                     *usermodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
		moreKeys                   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []usermodel.User
		wantErr bool
	}{
		{
			name:   "List users successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				userSearch:                 userSearch,
				filter:                     filterActive,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						filterActive.IsActive,
						filterActive.Role,
						userSearch).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						filterActive.IsActive,
						filterActive.Role,
						userSearch).
					WillReturnRows(rows)
			},
			want:    listUsers,
			wantErr: false,
		},
		{
			name:   "List users failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				userSearch:                 userSearch,
				filter:                     filterInactive,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filterInactive.SearchKey+"%",
						"%"+filterInactive.SearchKey+"%",
						"%"+filterInactive.SearchKey+"%",
						"%"+filterInactive.SearchKey+"%",
						"%"+filterInactive.SearchKey+"%",
						filterInactive.IsActive,
						filterInactive.Role,
						userSearch).
					WillReturnError(mockErr)
			},
			want:    listUsers,
			wantErr: true,
		},
		{
			name:   "List users failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				userSearch:                 userSearch,
				filter:                     filterActive,
				propertiesContainSearchKey: mockProperties,
				paging:                     paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						filterActive.IsActive,
						filterActive.Role,
						userSearch).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						"%"+filterActive.SearchKey+"%",
						filterActive.IsActive,
						filterActive.Role,
						userSearch).
					WillReturnError(mockErr)
			},
			want:    listUsers,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListUser(
				tt.args.ctx,
				tt.args.userSearch,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging,
				tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "ListUser() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListUser() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
