package importnotedetailstore

import (
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_FindListImportNoteDetail(t *testing.T) {
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

	conditions := map[string]interface{}{
		"importNoteId": "1",
	}

	expectedSql := "SELECT * FROM `ImportNoteDetail` WHERE `importNoteId` = ?"

	details := []importnotedetailmodel.ImportNoteDetail{
		{
			ImportNoteId: "1",
			IngredientId: "ingredient1",
			Price:        float32(20),
			AmountImport: 10,
			TotalUnit:    float32(200),
		},
		{
			ImportNoteId: "1",
			IngredientId: "ingredient2",
			Price:        float32(15),
			AmountImport: 5,
			TotalUnit:    float32(75),
		},
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
		want    []importnotedetailmodel.ImportNoteDetail
		wantErr bool
	}{
		{
			name: "Find list of import note details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs("1").
					WillReturnRows(
						sqlmock.NewRows([]string{"importNoteId", "ingredientId", "price", "amountImport", "totalUnit"}).
							AddRow("1", "ingredient1", float32(20), 10, float32(200)).
							AddRow("1", "ingredient2", float32(15), 5, float32(75)),
					)
			},
			want:    details,
			wantErr: false,
		},
		{
			name: "Find list of import note details failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs("1").
					WillReturnError(mockErr)
			},
			want:    details,
			wantErr: true,
		},
		{
			name: "Find list of import note details not found",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: conditions,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs("1").
					WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    details,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.FindListImportNoteDetail(tt.args.ctx, tt.args.conditions, tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "FindListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "FindListImportNoteDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
