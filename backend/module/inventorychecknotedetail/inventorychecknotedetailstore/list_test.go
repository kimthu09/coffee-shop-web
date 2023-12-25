package inventorychecknotedetailstore

import (
	"coffee_shop_management_backend/common/enum"
	"coffee_shop_management_backend/module/ingredient/ingredientmodel"
	"coffee_shop_management_backend/module/inventorychecknotedetail/inventorychecknotedetailmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_ListInventoryCheckNoteDetail(t *testing.T) {
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
		ctx                  context.Context
		inventoryCheckNoteId string
	}

	mockInventoryCheckNoteId := "note1"
	measureType := enum.Weight
	mockResult := []inventorychecknotedetailmodel.InventoryCheckNoteDetail{
		{
			InventoryCheckNoteId: mockInventoryCheckNoteId,
			IngredientId:         "ingredient1",
			Initial:              100,
			Difference:           10,
			Final:                110,
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:          "ingredient1",
				Name:        "Ingredient 1",
				MeasureType: &measureType,
			},
		},
		{
			InventoryCheckNoteId: mockInventoryCheckNoteId,
			IngredientId:         "ingredient2",
			Initial:              200,
			Difference:           20,
			Final:                220,
			Ingredient: ingredientmodel.SimpleIngredient{
				Id:          "ingredient2",
				Name:        "Ingredient 2",
				MeasureType: &measureType,
			},
		},
	}
	inventoryCheckNoteQuery := "SELECT * FROM `InventoryCheckNoteDetail` WHERE inventoryCheckNoteId = ?"
	ingredientQuery := "SELECT * FROM `Ingredient` WHERE `Ingredient`.`id` IN (?,?)"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []inventorychecknotedetailmodel.InventoryCheckNoteDetail
		wantErr bool
	}{
		{
			name: "List inventory check note details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: mockInventoryCheckNoteId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(inventoryCheckNoteQuery).
					WithArgs(mockInventoryCheckNoteId).
					WillReturnRows(sqlmock.NewRows([]string{"inventoryCheckNoteId", "ingredientId", "initial", "difference", "final"}).
						AddRow(mockResult[0].InventoryCheckNoteId, mockResult[0].IngredientId, mockResult[0].Initial, mockResult[0].Difference, mockResult[0].Final).
						AddRow(mockResult[1].InventoryCheckNoteId, mockResult[1].IngredientId, mockResult[1].Initial, mockResult[1].Difference, mockResult[1].Final))

				sqlDBMock.ExpectQuery(ingredientQuery).
					WithArgs(mockResult[0].IngredientId, mockResult[1].IngredientId).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "measureType"}).
						AddRow(mockResult[0].Ingredient.Id, mockResult[0].Ingredient.Name, mockResult[0].Ingredient.MeasureType).
						AddRow(mockResult[1].Ingredient.Id, mockResult[1].Ingredient.Name, mockResult[1].Ingredient.MeasureType))
			},
			want:    mockResult,
			wantErr: false,
		},
		{
			name: "Error handling",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:                  context.Background(),
				inventoryCheckNoteId: mockInventoryCheckNoteId,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(inventoryCheckNoteQuery).
					WithArgs(mockInventoryCheckNoteId).
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

			got, err := s.ListInventoryCheckNoteDetail(tt.args.ctx, tt.args.inventoryCheckNoteId)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInventoryCheckNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInventoryCheckNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got, "ListInventoryCheckNoteDetail() = %v, want %v", got, tt.want)
		})
	}
}
