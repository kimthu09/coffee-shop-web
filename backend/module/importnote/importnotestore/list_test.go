package importnotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/importnote/importnotemodel"
	"coffee_shop_management_backend/module/supplier/suppliermodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

func Test_sqlStore_ListImportNote(t *testing.T) {
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

	minPrice := float32(10)
	maxPrice := float32(90)
	filter := importnotemodel.Filter{
		SearchKey: "mockSearchKey",
		MinPrice:  &minPrice,
		MaxPrice:  &maxPrice,
		Status:    "Done",
	}
	properties := []string{"id", "supplierId", "createdBy", "closedBy"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	moreKeys := []string{"Supplier"}
	status := importnotemodel.Done
	validId := "012345678901"
	now := time.Now()
	listImportNote := []importnotemodel.ImportNote{
		{
			Id:         validId,
			SupplierId: validId,
			Supplier: suppliermodel.SimpleSupplier{
				Id:   validId,
				Name: "hello",
			},
			TotalPrice: 0,
			Status:     &status,
			CreatedBy:  validId,
			ClosedBy:   &validId,
			CreatedAt:  &now,
			ClosedAt:   &now,
		},
	}
	mockErr := errors.New(mock.Anything)

	rows := sqlmock.NewRows([]string{
		"id",
		"supplierId",
		"totalPrice",
		"status",
		"createdBy",
		"closedBy",
		"createdAt",
		"closedAt",
	})
	for _, importNote := range listImportNote {
		rows.AddRow(
			importNote.Id,
			importNote.SupplierId,
			importNote.TotalPrice,
			importNote.Status,
			importNote.CreatedBy,
			importNote.ClosedBy,
			importNote.CreatedAt,
			importNote.ClosedAt)
	}

	queryString := "SELECT * FROM `ImportNote` " +
		"WHERE (id LIKE ? OR supplierId LIKE ? OR createdBy LIKE ? OR closedBy LIKE ?) " +
		"AND status = ? " +
		"AND totalPrice >= ? " +
		"AND totalPrice <= ? " +
		"ORDER BY createdAt desc LIMIT " + strconv.FormatInt(paging.Limit, 10)

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listImportNote))
	queryStringCount := "SELECT count(*) FROM `ImportNote` " +
		"WHERE (id LIKE ? OR supplierId LIKE ? OR createdBy LIKE ? OR closedBy LIKE ?) " +
		"AND status = ? " +
		"AND totalPrice >= ? " +
		"AND totalPrice <= ?"

	supplierRow := sqlmock.NewRows([]string{
		"id",
		"name",
	})
	supplierRow.AddRow(
		listImportNote[0].Supplier.Id,
		listImportNote[0].Supplier.Name)

	querySupplier := "SELECT * FROM `Supplier` WHERE `Supplier`.`id` = ?"

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx                        context.Context
		filter                     *importnotemodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
		moreKeys                   []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []importnotemodel.ImportNote
		wantErr bool
	}{
		{
			name:   "List import note failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: properties,
				paging:                     &paging,
				moreKeys:                   moreKeys,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.Status,
						*filter.MinPrice,
						*filter.MaxPrice,
					).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "List import note failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: properties,
				paging:                     &paging,
				moreKeys:                   moreKeys,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.Status,
						*filter.MinPrice,
						*filter.MaxPrice,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.Status,
						*filter.MinPrice,
						*filter.MaxPrice,
					).
					WillReturnError(mockErr)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:   "List import note successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: properties,
				paging:                     &paging,
				moreKeys:                   moreKeys,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.Status,
						*filter.MinPrice,
						*filter.MaxPrice,
					).
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						"%"+filter.SearchKey+"%",
						filter.Status,
						*filter.MinPrice,
						*filter.MaxPrice,
					).
					WillReturnRows(rows)

				mockSqlDB.
					ExpectQuery(querySupplier).
					WithArgs(validId).
					WillReturnRows(supplierRow)

			},
			want:    listImportNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListImportNote(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging,
				tt.args.moreKeys...)

			if tt.wantErr {
				assert.NotNil(t, err, "ListIngredient() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListIngredient() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListIngredient() = %v, want %v", got, tt.want)
			}
		})
	}
}
