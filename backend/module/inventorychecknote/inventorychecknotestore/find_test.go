package inventorychecknotestore

import (
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

func Test_sqlStore_FindInventoryCheckNote(t *testing.T) {
	sqlDB, sqlDBMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
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
		ctx        context.Context
		conditions map[string]interface{}
		moreKeys   []string
	}

	inventoryCheckNoteId := mock.Anything
	mockConditions := map[string]interface{}{
		"id": inventoryCheckNoteId,
	}
	mockData := &inventorychecknotemodel.InventoryCheckNote{
		Id:                "someID",
		AmountDifferent:   10,
		AmountAfterAdjust: 20,
		CreatedBy:         "user123",
		CreatedAt:         &time.Time{},
	}
	expectedQuery := "SELECT * FROM `InventoryCheckNote` WHERE `id` = ? ORDER BY `InventoryCheckNote`.`id` LIMIT 1"
	mockErr := errors.New(mock.Anything)

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		want    *inventorychecknotemodel.InventoryCheckNote
		wantErr bool
	}{
		{
			name: "Find inventory check note successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "amountDifferent", "amountAfterAdjust", "createdBy", "createdAt"}).
					AddRow(mockData.Id, mockData.AmountDifferent, mockData.AmountAfterAdjust, mockData.CreatedBy, mockData.CreatedAt)
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(mockConditions["id"]).WillReturnRows(rows)
			},
			want:    mockData,
			wantErr: false,
		},
		{
			name: "Find inventory check note failed because can not find inventory check note",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(mockConditions["id"]).WillReturnError(gorm.ErrRecordNotFound)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Find inventory check note failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:        context.Background(),
				conditions: mockConditions,
			},
			mock: func() {
				sqlDBMock.ExpectQuery(expectedQuery).
					WithArgs(mockConditions["id"]).WillReturnError(mockErr)
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

			got, err := s.FindInventoryCheckNote(tt.args.ctx, tt.args.conditions)

			if tt.wantErr {
				assert.NotNil(t, err, "FindInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "FindInventoryCheckNote() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.Equal(t, tt.want, got, "FindInventoryCheckNote() = %v, want %v", got, tt.want)
		})
	}
}
