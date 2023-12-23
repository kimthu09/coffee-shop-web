package rolemodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/rolefeature/rolefeaturemodel"
	"errors"
)

type Role struct {
	Id           string           `json:"id" gorm:"column:id;"`
	Name         string           `json:"name" gorm:"column:name;"`
	RoleFeatures ListRoleFeatures `json:"features,omitempty"`
}

func (*Role) TableName() string {
	return common.TableRole
}

type ListRoleFeatures []rolefeaturemodel.RoleFeature

func (*ListRoleFeatures) TableName() string {
	return common.TableRoleFeature
}

var (
	ErrRoleNameEmpty = common.NewCustomError(
		errors.New("name of role is empty"),
		"Tên của vai trò đang trống",
		"ErrRoleNameEmpty",
	)
	ErrRoleFeaturesEmpty = common.NewCustomError(
		errors.New("features of role is empty"),
		"Danh sách chức năng của vai trò đang trống",
		"ErrRoleFeaturesEmpty",
	)
	ErrRoleFeatureInvalid = common.NewCustomError(
		errors.New("features of role is invalid"),
		"Chức năng không hợp lệ",
		"ErrRoleFeatureInvalid",
	)
	ErrRoleNameDuplicate = common.ErrDuplicateKey(
		errors.New("Tên của vai trò đã tồn tại"),
	)
	ErrRoleCreateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền tạo vai trò mới"),
	)
	ErrRoleUpdateNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền chỉnh sửa thông tin vai trò"),
	)
	ErrRoleViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem vai trò"),
	)
)
