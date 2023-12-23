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
		"Mã của phiếu hủy không hợp lệ",
		"ErrExportNoteIdInvalid",
	)
	ErrExportNoteReasonEmpty = common.NewCustomError(
		errors.New("export reason is empty"),
		"Lý do hủy không hợp lệ",
		"ErrExportNoteReasonEmpty",
	)
	ErrExportNoteDetailsEmpty = common.NewCustomError(
		errors.New("list export note details are empty"),
		"Danh sách các nguyên vật liệu muốn hủy đang trống",
		"ErrExportNoteDetailsEmpty",
	)
	ErrExportNoteExistDuplicateIngredient = common.NewCustomError(
		errors.New("exist duplicate ingredient"),
		"Tồn tại 2 nguyên vật liệu trùng nhau",
		"ErrExportNoteExistDuplicateIngredient",
	)
	ErrExportNoteAmountExportIsOverTheStock = common.NewCustomError(
		errors.New("amount export is over stock"),
		"Lượng muốn xuất vượt quá lượng trong kho",
		"ErrExportNoteAmountExportIsOverTheStock",
	)
	ErrExportNoteCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo phiếu xuất"),
	)
	ErrExportNoteViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem phiếu xuất"),
	)
)
