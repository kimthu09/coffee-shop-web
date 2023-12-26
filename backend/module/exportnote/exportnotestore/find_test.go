package exportnotestore

import (
	"coffee_shop_management_backend/module/exportnote/exportnotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_sqlStore_FindExportNote(t *testing.T) {
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
		moreKeys   []string
	}

	reason := exportnotemodel.Damaged
	exportNoteId := "ExportNote001"
	mockData := &exportnotemodel.ExportNote{
		Id:        exportNoteId,
		Reason:    &reason,
		CreatedAt: &time.Time{},
		CreatedBy: "User001",
	}

	expectedQuery := "SELECT * FROM `ExportNote` WHERE `id` = ? ORDER BY `ExportNote`.`id` LIMIT 1"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *exportnotemodel.ExportNote
		wantErr bool
	}{
		{
			name: "Find export note successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"id": exportNoteId,
				},
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "reason", "createdAt", "createdBy"}).
					AddRow(mockData.Id, mockData.Reason, mockData.CreatedAt, mockData.CreatedBy)

				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(exportNoteId).
					WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Find export note failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx: context.Background(),
				conditions: map[string]interface{}{
					"id": exportNoteId,
				},
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(exportNoteId).
					WillReturnError(mockErr)
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

			got, err := s.FindExportNote(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindIngredient() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindIngredient() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindIngredient() got = %v, want %v", got, tt.want)
			}
		})
	}
}
