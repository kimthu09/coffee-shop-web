package importnotestore

import (
	"coffee_shop_management_backend/module/importnote/importnotemodel"
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

func Test_sqlStore_FindImportNote(t *testing.T) {
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
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	validId := "123"
	expectedSql := "SELECT * FROM `ImportNote` " +
		"WHERE `id` = ? " +
		"ORDER BY `ImportNote`.`id` LIMIT 1"
	status := importnotemodel.Done
	now := time.Now()
	importNote := importnotemodel.ImportNote{
		Id:         validId,
		SupplierId: mock.Anything,
		TotalPrice: 0,
		Status:     &status,
		CreatedBy:  mock.Anything,
		ClosedBy:   &validId,
		CreatedAt:  &now,
		ClosedAt:   &now,
	}
	importNoteRows := sqlmock.NewRows([]string{
		"id",
		"supplierId",
		"totalPrice",
		"status",
		"createdBy",
		"closedBy",
		"createdAt",
		"closedAt",
	})
	importNoteRows.AddRow(importNote.Id,
		importNote.SupplierId,
		importNote.TotalPrice,
		importNote.Status,
		importNote.CreatedBy,
		importNote.ClosedBy,
		importNote.CreatedAt,
		importNote.ClosedAt)
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name: "Find import note successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: map[string]interface{}{"id": validId},
				moreKeys:   nil,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(validId).
					WillReturnRows(importNoteRows)
			},
			want:    &importNote,
			wantErr: false,
		},
		{
			name: "Find import note failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: map[string]interface{}{"id": validId},
				moreKeys:   nil,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(validId).
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

			got, err := s.FindImportNote(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindImportNote() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindImportNote() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}
