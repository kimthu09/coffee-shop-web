package exportnotemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/exportnotedetail/exportnotedetailmodel"
	"coffee_shop_management_backend/module/user/usermodel"
	"errors"
	"time"
)

type ExportNote struct {
	Id            string                                   `json:"id" gorm:"column:id;"`
	Reason        *ExportReason                            `json:"reason" gorm:"column:reason;"`
	CreatedAt     *time.Time                               `json:"createdAt" gorm:"column:createdAt;"`
	CreatedBy     string                                   `json:"-" gorm:"column:createdBy;"`
	CreatedByUser usermodel.SimpleUser                     `json:"createdBy" gorm:"foreignKey:CreatedBy;references:Id"`
	Details       []exportnotedetailmodel.ExportNoteDetail `json:"details,omitempty"`
}

func (*ExportNote) TableName() string {
	return common.TableExportNote
}

var (
	ErrExportNoteIdInvalid = common.NewCustomError(
		errors.New("id of export note is invalid"),
		"id of export note is invalid",
		"ErrExportNoteIdInvalid",
	)
	ErrExportNoteReasonEmpty = common.NewCustomError(
		errors.New("export reason is empty"),
		"export reason is empty",
		"ErrExportNoteReasonEmpty",
	)
	ErrExportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list export note details are empty"),
		"list export note details are empty",
		"ErrExportNoteDetailsEmpty",
	)
	ErrExportNoteExistDuplicateIngredient = common.NewCustomError(
		errors.New("exist duplicate ingredient"),
		"exist duplicate ingredient",
		"ErrExportNoteExistDuplicateIngredient",
	)
	ErrExportNoteAmountExportIsOverTheStock = common.NewCustomError(
		errors.New("amount export is over stock"),
		"amount export is over stock",
		"ErrExportNoteAmountExportIsOverTheStock",
	)
	ErrExportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to create export note"),
	)
	ErrExportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("you have no permission to view export note"),
	)
)
