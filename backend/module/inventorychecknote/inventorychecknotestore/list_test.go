package inventorychecknotestore

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/inventorychecknote/inventorychecknotemodel"
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func Test_sqlStore_ListInventoryCheckNote(t *testing.T) {
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
		ctx                        context.Context
		filter                     *inventorychecknotemodel.Filter
		propertiesContainSearchKey []string
		paging                     *common.Paging
	}
	timeInt := int64(123)
	timeDate := time.Unix(123, 0)
	createdBy := "user001"
	listInventoryCheckNote := []inventorychecknotemodel.InventoryCheckNote{
		{
			Id:                "123",
			AmountDifferent:   10,
			AmountAfterAdjust: 120,
			CreatedBy:         "user001",
			CreatedAt:         &timeDate,
		},
	}
	filter := inventorychecknotemodel.Filter{
		SearchKey:         "mockSearchKey",
		DateFromCreatedAt: &timeInt,
		DateToCreatedAt:   &timeInt,
		CreatedBy:         &createdBy,
	}
	mockProperties := []string{"id"}
	paging := common.Paging{
		Page:  1,
		Limit: 10,
	}
	rows := sqlmock.NewRows([]string{
		"id",
		"amountDifferent",
		"amountAfterAdjust",
		"createdBy",
		"createdAt",
	})
	for _, exportNote := range listInventoryCheckNote {
		rows.AddRow(
			exportNote.Id,
			exportNote.AmountDifferent,
			exportNote.AmountAfterAdjust,
			exportNote.CreatedBy,
			exportNote.CreatedAt,
		)
	}

	queryString :=
		"SELECT `InventoryCheckNote`.`id`,`InventoryCheckNote`.`amountDifferent`,`InventoryCheckNote`.`amountAfterAdjust`,`InventoryCheckNote`.`createdBy`,`InventoryCheckNote`.`createdAt` FROM `InventoryCheckNote` " +
			"JOIN MUser ON InventoryCheckNote.createdBy = MUser.id " +
			"WHERE id LIKE ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND MUser.id = ? " +
			"ORDER BY createdAt desc LIMIT 10"

	countRows := sqlmock.NewRows([]string{
		"count",
	})
	countRows.AddRow(len(listInventoryCheckNote))
	queryStringCount :=
		"SELECT count(*) FROM `InventoryCheckNote` " +
			"JOIN MUser ON InventoryCheckNote.createdBy = MUser.id " +
			"WHERE id LIKE ? " +
			"AND createdAt >= ? " +
			"AND createdAt <= ? " +
			"AND MUser.id = ?"

	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    []inventorychecknotemodel.InventoryCheckNote
		wantErr bool
	}{
		{
			name:   "List inventory check note failed because can not get number of rows from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						timeDate,
						timeDate,
						"user001").
					WillReturnRows(countRows)
			},
			want:    listInventoryCheckNote,
			wantErr: true,
		},
		{
			name:   "List inventory check note failed because can not get list from database",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						timeDate,
						timeDate,
						"user001").
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						timeDate,
						timeDate,
						"user001").
					WillReturnError(mockErr)
			},
			want:    listInventoryCheckNote,
			wantErr: true,
		},
		{
			name:   "List inventory check note successfully",
			fields: fields{db: gormDB},
			args: args{
				ctx:                        context.Background(),
				filter:                     &filter,
				propertiesContainSearchKey: mockProperties,
				paging:                     &paging,
			},
			mock: func() {
				mockSqlDB.
					ExpectQuery(queryStringCount).
					WithArgs(
						"%"+filter.SearchKey+"%",
						timeDate,
						timeDate,
						"user001").
					WillReturnRows(countRows)

				mockSqlDB.
					ExpectQuery(queryString).
					WithArgs(
						"%"+filter.SearchKey+"%",
						timeDate,
						timeDate,
						"user001").
					WillReturnRows(rows)
			},
			want:    listInventoryCheckNote,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlStore{
				db: tt.fields.db,
			}

			tt.mock()

			got, err := s.ListInventoryCheckNote(
				tt.args.ctx,
				tt.args.filter,
				tt.args.propertiesContainSearchKey,
				tt.args.paging)

			if tt.wantErr {
				assert.NotNil(t, err, "ListInventoryCheckNote() err = %v, wantErr = %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "ListInventoryCheckNote() err = %v, wantErr = %v", err, tt.wantErr)
				assert.Equal(t, tt.want, got, "ListInventoryCheckNote() = %v, want %v", got, tt.want)
			}
		})
	}
}
