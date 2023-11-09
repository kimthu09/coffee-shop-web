package exportnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
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
			ExpiryDate:   "09/11/2023",
			AmountExport: 100,
		},
		{
			ExportNoteId: exportNoteId,
			IngredientId: "Ing002",
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:   "Ing002",
				Name: "Nguyen vat lieu 002",
			},
			ExpiryDate:   "10/11/2023",
			AmountExport: 100,
		},
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	mockErr := errors.New(mock.Anything)

	rows := sqlmock.NewRows([]string{
		"exportNoteId",
		"ingredientId",
		"expiryDate",
		"amountExport",
	})
	for _, detail := range exportNoteDetails {
		rows.AddRow(
			detail.ExportNoteId,
			detail.IngredientId,
			detail.ExpiryDate,
			detail.AmountExport)
	}
	queryString := fmt.Sprintf(
		"SELECT * FROM `ExportNoteDetail` WHERE exportNoteId = ? LIMIT %v",
		strconv.FormatInt(paging.Limit, 10))

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(exportNoteDetails))
	queryStringCount :=
		"SELECT count(*) FROM `ExportNoteDetail` WHERE exportNoteId = ?"

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
			name:   "List export note detail failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(exportNoteId).
					WillReturnError(mockErr)
			},
			want:    exportNoteDetails,
			wantErr: true,
		},
		{
			name:   "List export note detail failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				exportNoteId: exportNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(exportNoteId).
					WillReturnRows(countRows)

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
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(exportNoteId).
					WillReturnRows(countRows)

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

			got, err := s.ListExportNoteDetail(tt.args.ctx, tt.args.exportNoteId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListExportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListExportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListExportNoteDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
