package exportnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListExportNoteDetail(t *testing.T) {
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
		ctx          context.Context
		exportNoteId string
		paging       *common.Paging
	}

	exportNoteId := mock.Anything
	exportNoteDetails := []exportnotedetailmodel.ExportNoteDetail{
		{
			ExportNoteId: exportNoteId,
			IngredientId: "Ing001",
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:   "Ing001",
				Name: "Nguyen vat lieu 001",
			},
			AmountExport: 100,
		},
		{
			ExportNoteId: exportNoteId,
			IngredientId: "Ing002",
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:   "Ing002",
				Name: "Nguyen vat lieu 002",
			},
			AmountExport: 100,
		},
	}
	mockErr := errors.New(mock.Anything)

	rows := sqlmock.NewRows([]string{
		"exportNoteId",
		"ingredientId",
		"amountExport",
	})
	for _, detail := range exportNoteDetails {
		rows.AddRow(
			detail.ExportNoteId,
			detail.IngredientId,
			detail.AmountExport)
	}
	queryString :=
		"SELECT * FROM `ExportNoteDetail` WHERE exportNoteId = ?"

	queryIngredient := "SELECT * FROM `Ingredient` WHERE `Ingredient`.`id` IN (?,?) ORDER BY Ingredient.name"
	ingredientRow := sqlmock.NewRows([]string{
		"id",
		"name",
	})
	for _, v := range exportNoteDetails {
		ingredientRow.AddRow(v.Ingredient.Id, v.Ingredient.Name)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []exportnotedetailmodel.ExportNoteDetail
		wantErr bool
	}{
		{
			name:   "List export note detail failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(exportNoteId).
					WillReturnError(mockErr)
			},
			want:    exportNoteDetails,
			wantErr: true,
		},
		{
			name:   "List export note detail successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(exportNoteId).
					WillReturnRows(rows)

				mockSqlDB.
					ExpectQuery(queryIngredient).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(ingredientRow)

			},
			want:    exportNoteDetails,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListExportNoteDetail(tt.args.ctx, tt.args.exportNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "ListExportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListExportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListExportNoteDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
