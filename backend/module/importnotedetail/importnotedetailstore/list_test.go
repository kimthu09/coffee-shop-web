package importnotedetailstore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
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

func Test_sqlStore_ListImportNoteDetail(t *testing.T) {
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

	importNoteID := "1"

	expectedSql := "SELECT * FROM `ImportNoteDetail` WHERE importNoteId = ?"
	expectedIngredientPreloadSql := "SELECT * FROM `Ingredient` WHERE `Ingredient`.`id` IN (?,?) ORDER BY Ingredient.name"

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
			Price:        15,
			AmountImport: 5,
			TotalUnit:    float32(75),
		},
	}

	mockErr := errors.New(mock.Anything)
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx          context.Context
		importNoteId string
		paging       *common.Paging
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
			name: "List import note details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteID,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(importNoteID).
					WillReturnRows(
						sqlmock.NewRows([]string{"importNoteId", "ingredientId", "price", "amountImport", "totalUnit"}).
							AddRow("1", "ingredient1", float32(20), 10, float32(200)).
							AddRow("1", "ingredient2", float32(15), 5, float32(75)),
					)

				sqlDBMock.
					ExpectQuery(expectedIngredientPreloadSql).
					WithArgs("ingredient1", "ingredient2").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name"}).
							AddRow("ingredient1", "Ingredient 1").
							AddRow("ingredient2", "Ingredient 2"),
					)

				details[0].Ingredient = ingredientmodel.SimpleIngredient{
					Id:   "ingredient1",
					Name: "Ingredient 1",
				}

				details[1].Ingredient = ingredientmodel.SimpleIngredient{
					Id:   "ingredient2",
					Name: "Ingredient 2",
				}
			},
			want:    details,
			wantErr: false,
		},
		{
			name: "List import note details failed because can not get list import note from database",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteID,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(importNoteID).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "List import note details failed because can not preload ingredient",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:          context.Background(),
				importNoteId: importNoteID,
			},
			mock: func() {
				sqlDBMock.
					ExpectQuery(expectedSql).
					WithArgs(importNoteID).
					WillReturnRows(
						sqlmock.NewRows([]string{"importNoteId", "ingredientId", "price", "amountImport", "totalUnit"}).
							AddRow("1", "ingredient1", float32(20), 10, float32(200)).
							AddRow("1", "ingredient2", float32(15), 5, float32(75)),
					)

				sqlDBMock.
					ExpectQuery(expectedIngredientPreloadSql).
					WithArgs("ingredient1", "ingredient2").
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

			got, err := s.ListImportNoteDetail(tt.args.ctx, tt.args.importNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "ListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListImportNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListImportNoteDetail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
