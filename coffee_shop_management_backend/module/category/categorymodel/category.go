package categorymodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type SimpleCategory struct {
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}

type Category struct {
	Id            string `json:"id" gorm:"column:id;"`
	Name          string `json:"name" gorm:"column:name;"`
	Description   string `json:"description" gorm:"column:description;"`
	AmountProduct int    `json:"amountProduct" gorm:"column:amountProduct;"`
}

func (*Category) TableName() string {
	return common.TableCategory
}

var (
	ErrCategoryNameEmpty = common.NewCustomError(
		errors.New("name of category is empty"),
		"Tên của danh mục đang trống",
		"ErrCategoryNameEmpty",
	)
	ErrCategoryIdDuplicate = common.ErrDuplicateKey(
		errors.New("Đã tồn tại danh mục trong hệ thống"),
	)
	ErrCategoryNameDuplicate = common.ErrDuplicateKey(
		errors.New("Đã tồn tại danh mục có tên này trong hệ thống"),
	)
	ErrCategoryCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo danh mục mới"),
	)
	ErrCategoryUpdateInfoNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin danh mục"),
	)
	ErrCategoryViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem danh mục"),
	)
)
