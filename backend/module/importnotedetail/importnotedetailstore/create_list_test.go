package importnotedetailstore

import (
	"coffee_shop_management_backend/module/importnotedetail/importnotedetailmodel"
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

	importNoteDetails := []importnotedetailmodel.ImportNoteDetailCreate{
		{ImportNoteId: "1", IngredientId: "ingredient1", Price: float32(20), AmountImport: 10, TotalUnit: float32(200)},
		{ImportNoteId: "1", IngredientId: "ingredient2", Price: float32(15), AmountImport: 5, TotalUnit: float32(75)},
	}

	expectedSql := "INSERT INTO `ImportNoteDetail` (`importNoteId`,`ingredientId`,`price`,`amountImport`,`totalUnit`) VALUES (?,?,?,?,?),(?,?,?,?,?)"

	mockErr := errors.New(mock.Anything)

	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx  context.Context
		data []importnotedetailmodel.ImportNoteDetailCreate
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Create list of import note details successfully",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: importNoteDetails,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						"1", "ingredient1", float32(20), 10, float32(200),
						"1", "ingredient2", float32(15), 5, float32(75),
					).
					WillReturnResult(sqlmock.NewResult(1, 2))
				sqlDBMock.ExpectCommit()
			},
			wantErr: false,
		},
		{
			name: "Create list of import note details failed",
			fields: fields{
				db: gormDB,
			},
			args: args{
				ctx:  context.Background(),
				data: importNoteDetails,
			},
			mock: func() {
				sqlDBMock.ExpectBegin()
				sqlDBMock.
					ExpectExec(expectedSql).
					WithArgs(
						"1", "ingredient1", float32(20), 10, float32(200),
						"1", "ingredient2", float32(15), 5, float32(75),
					).
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

			err := s.CreateListImportNoteDetail(tt.args.ctx, tt.args.data)

			if tt.wantErr {
				assert.NotNil(t, err, "CreateListRoleFeatureDetail() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				assert.Nil(t, err, "CreateListRoleFeatureDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
