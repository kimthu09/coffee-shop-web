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

func Test_sqlStore_FindFeature(t *testing.T) {
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

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		id  string
	}

	featureId := mock.Anything
	feature := featuremodel.Feature{
		Id:          featureId,
		Description: mock.Anything,
	}
	featureRows := sqlmock.NewRows([]string{
		"id",
		"description",
	})
	featureRows.AddRow(
		feature.Id,
		feature.Description)
	mockErr :=
		errors.New(mock.Anything)
	expectedSql := "SELECT * FROM `Feature` " +
		"WHERE id = ? " +
		"ORDER BY `Feature`.`id` " +
		"LIMIT 1"

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *featuremodel.Feature
		wantErr bool
	}{
		{
			name: "Find feature failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				id:  featureId,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(featureId).
					WillReturnError(mockErr)
			},
			want:    &feature,
			wantErr: true,
		},
		{
			name: "Find feature successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				id:  featureId,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(featureId).
					WillReturnRows(featureRows)
			},
			want:    &feature,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindFeature(tt.args.ctx, tt.args.id)

			if tt.wantErr {
				assert.NotNil(t, err, "FindFeature() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindFeature() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindFeature() = %v, want %v", got, tt.want)
			}
		})
	}
}
