package featurestore

import (
	"coffee_shop_management_backend/module/feature/featuremodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListFeature(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
	}

	features := []featuremodel.Feature{
		{
			Id:          mock.Anything,
			Description: mock.Anything,
		},
		{
			Id:          mock.Anything,
			Description: mock.Anything,
		},
		{
			Id:          mock.Anything,
			Description: mock.Anything,
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"description",
	})
	for _, v := range features {
		rows.AddRow(v.Id, v.Description)
	}

	queryString := "SELECT * FROM `Feature`"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []featuremodel.Feature
		wantErr bool
	}{
		{
			name:   "List features failed",
			fields: fields{db: gormDB},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryString).
					WillReturnError(mockErr)
			},
			want:    features,
			wantErr: true,
		},
		{
			name:   "List features successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx: context.Background(),
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryString).
					WillReturnRows(rows)
			},
			want:    features,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListFeature(tt.args.ctx)

			if tt.wantErr {
				assert.NotNil(t, err, "ListFeature() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListFeature() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}
