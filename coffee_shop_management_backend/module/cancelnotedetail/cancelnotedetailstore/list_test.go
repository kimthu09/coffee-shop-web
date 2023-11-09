package cancelnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/cancelnotedetail/cancelnotedetailmodel"
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

func Test_sqlStore_ListCancelNoteDetail(t *testing.T) {
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
		cancelNoteId string
		paging       *common.Paging
	}

	cancelNoteId := mock.Anything
	cancelNoteDetails := []cancelnotedetailmodel.CancelNoteDetail{
		{
			CancelNoteId: cancelNoteId,
			IngredientId: "Ing001",
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:   "Ing001",
				Name: "Nguyen vat lieu 001",
			},
			ExpiryDate:   "09/11/2023",
			Reason:       cancelnotedetailmodel.Damaged,
			AmountCancel: 100,
		},
		{
			CancelNoteId: cancelNoteId,
			IngredientId: "Ing002",
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:   "Ing002",
				Name: "Nguyen vat lieu 002",
			},
			ExpiryDate:   "10/11/2023",
			Reason:       cancelnotedetailmodel.Damaged,
			AmountCancel: 100,
		},
	}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	mockErr := errors.New(mock.Anything)

	rows := sqlmock.NewRows([]string{
		"cancelNoteId",
		"ingredientId",
		"expiryDate",
		"reason",
		"amountCancel",
	})
	for _, detail := range cancelNoteDetails {
		rows.AddRow(
			detail.CancelNoteId,
			detail.IngredientId,
			detail.ExpiryDate,
			[]byte(detail.Reason.String()),
			detail.AmountCancel)
	}
	queryString := fmt.Sprintf(
		"SELECT * FROM `CancelNoteDetail` WHERE cancelNoteId = ? LIMIT %v",
		strconv.FormatInt(paging.Limit, 10))

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(cancelNoteDetails))
	queryStringCount :=
		"SELECT count(*) FROM `CancelNoteDetail` WHERE cancelNoteId = ?"

	queryIngredient := "SELECT * FROM `Ingredient` WHERE `Ingredient`.`id` IN (?,?) ORDER BY Ingredient.name"
	ingredientRow := sqlmock.NewRows([]string{
		"id",
		"name",
	})
	for _, v := range cancelNoteDetails {
		ingredientRow.AddRow(v.Ingredient.Id, v.Ingredient.Name)
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []cancelnotedetailmodel.CancelNoteDetail
		wantErr bool
	}{
		{
			name:   "List cancel note detail failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: cancelNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(cancelNoteId).
					WillReturnError(mockErr)
			},
			want:    cancelNoteDetails,
			wantErr: true,
		},
		{
			name:   "List cancel note detail failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: cancelNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(cancelNoteId).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(cancelNoteId).
					WillReturnError(mockErr)
			},
			want:    cancelNoteDetails,
			wantErr: true,
		},
		{
			name:   "List cancel note detail successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:          context.Background(),
				cancelNoteId: cancelNoteId,
				paging:       &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(cancelNoteId).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(cancelNoteId).
					WillReturnRows(rows)

				mockSqlDB.
					ExpectQuery(queryIngredient).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(ingredientRow)

			},
			want:    cancelNoteDetails,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListCancelNoteDetail(tt.args.ctx, tt.args.cancelNoteId, tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListCancelNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListCancelNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListCancelNoteDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
