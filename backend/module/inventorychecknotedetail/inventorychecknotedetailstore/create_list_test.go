package inventorychecknotedetailstore

import (
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

func Test_sqlStore_CreateListInventoryCheckNoteDetail(t *testing.T) {
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
		ctx  context.Context
		data []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate
	}

	mockData := []inventorychecknotedetailmodel.InventoryCheckNoteDetailCreate{
		{
			InventoryCheckNoteId: "note1",
			IngredientId:         "ingredient1",
			Initial:              100,
			Difference:           10,
			Final:                110,
		},
		{
			InventoryCheckNoteId: "note2",
			IngredientId:         "ingredient2",
			Initial:              200,
			Difference:           20,
			Final:                220,
		},
	}
	mockErr := errors.New(mock.Anything)
	expectedQuery := "INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`,`ingredientId`,`initial`,`difference`,`final`) VALUES (?,?,?,?,?),(?,?,?,?,?)"
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create list of inventory check note details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WithArgs(
						mockData[0].InventoryCheckNoteId,
						mockData[0].IngredientId,
						mockData[0].Initial,
						mockData[0].Difference,
						mockData[0].Final,
						mockData[1].InventoryCheckNoteId,
						mockData[1].IngredientId,
						mockData[1].Initial,
						mockData[1].Difference,
						mockData[1].Final,
					).
					WillReturnResult(
						sqlmock.NewResult(1, 2))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create list of inventory check note details failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: mockData,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.ExpectExec(expectedQuery).
					WillReturnError(mockErr)
				sqlDBMock.ExpectRollback()
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			err := s.CreateListInventoryCheckNoteDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListInventoryCheckNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListInventoryCheckNoteDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
