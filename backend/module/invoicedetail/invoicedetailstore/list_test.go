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

func Test_sqlStore_ListInvoiceDetail(t *testing.T) {
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

	invoiceId := "invoice123"
	sizeName := "Medium"
	amount := 2
	unitPrice := 10
	description := "Description 1"

	details := []invoicedetailmodel.InvoiceDetail{
		{
			InvoiceId: invoiceId,
			FoodId:    "food123",
			Food: invoicedetailmodel.SimpleFood{
				Id:   "food123",
				Name: "Food 1",
			},
			SizeName:    sizeName,
			Amount:      amount,
			UnitPrice:   unitPrice,
			Description: description,
		},
	}
	expectedQuery :=
		"SELECT * FROM `InvoiceDetail` WHERE invoiceId = ?"
	expectedFoodQuery := "SELECT * FROM `Food` WHERE `Food`.`id` = ? ORDER BY Food.name desc"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx       context.Context
		invoiceId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []invoicedetailmodel.InvoiceDetail
		wantErr bool
	}{
		{
			name:   "List invoice detail successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedQuery).
					WithArgs(invoiceId).
					WillReturnRows(
						sqlmock.NewRows([]string{"invoiceId", "foodId", "sizeName", "amount", "unitPrice", "description"}).
							AddRow(invoiceId, "food123", sizeName, amount, unitPrice, description),
					)

				mockSqlDB.
					ExpectQuery(expectedFoodQuery).
					WithArgs("food123").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name"}).
							AddRow("food123", "Food 1"),
					)
			},
			want:    details,
			wantErr: false,
		},
		{
			name:   "List invoice detail failed because can not get invoice",
			fields: fields{db: gormDB},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedQuery).
					WithArgs(invoiceId).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "List invoice detail failed because can not get details of invoice",
			fields: fields{db: gormDB},
			args: args{
				ctx:       context.Background(),
				invoiceId: invoiceId,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(expectedQuery).
					WithArgs(invoiceId).
					WillReturnRows(
						sqlmock.NewRows([]string{"invoiceId", "foodId", "sizeName", "amount", "unitPrice", "description"}).
							AddRow(invoiceId, "food123", sizeName, amount, unitPrice, description),
					)

				mockSqlDB.
					ExpectQuery(expectedFoodQuery).
					WithArgs("food123").
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

			got, err := s.ListInvoiceDetail(tt.args.ctx, tt.args.invoiceId)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInvoiceDetail() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInvoiceDetail() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListInvoiceDetail() = %v, want %v", got, tt.want)
			}
		})
	}
}
