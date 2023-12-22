package invoicedetailstore

import (
	"coffee_shop_management_backend/module/invoicedetail/invoicedetailmodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func Test_sqlStore_CreateListImportNoteDetail(t *testing.T) {
	sqlDB, mockSqlDB, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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

	details := []invoicedetailmodel.InvoiceDetailCreate{
		{
			InvoiceId: "invoice123",
			FoodId:    "food123",
			FoodName:  "Food 1",
			SizeId:    "size123",
			SizeName:  "Medium",
			Toppings: &invoicedetailmodel.InvoiceDetailToppings{
				{
					Id:    "Topping1",
					Name:  "Topping 1",
					Price: 10000,
				},
			},
			Amount:      2,
			UnitPrice:   10,
			Description: "Description 1",
		},
		{
			InvoiceId: "invoice123",
			FoodId:    "food456",
			FoodName:  "Food 2",
			SizeId:    "size456",
			SizeName:  "Large",
			Toppings: &invoicedetailmodel.InvoiceDetailToppings{
				{
					Id:    "Topping2",
					Name:  "Topping 2",
					Price: 10000,
				},
			},
			Amount:      3,
			UnitPrice:   15,
			Description: "Description 2",
		},
	}
	expectedQuery := "INSERT INTO `InvoiceDetail` (`invoiceId`,`foodId`,`foodName`,`sizeName`,`toppings`,`amount`,`unitPrice`,`description`) " +
		"VALUES (?,?,?,?,?,?,?,?),(?,?,?,?,?,?,?,?)"
	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data []invoicedetailmodel.InvoiceDetailCreate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name:   "Create list import note detail successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: details,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WillReturnResult(sqlmock.NewResult(0, int64(len(details))))
				mockSqlDB.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name:   "Create list import note detail failed",
			fields: fields{db: gormDB},
			args: args{
				ctx:  context.Background(),
				data: details,
			},
			mock: func() {
				mockSqlDB.ExpectBegin()
				mockSqlDB.
					ExpectExec(expectedQuery).
					WillReturnError(mockErr)
				mockSqlDB.ExpectRollback()
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

			err := s.CreateListImportNoteDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListImportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListImportNoteDetail() err = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
