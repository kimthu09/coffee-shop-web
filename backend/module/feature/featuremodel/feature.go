package featuremodel

import (
	"coffee_shop_management_backend/common"
	"errors"
)

type Feature struct {
	Id          string `json:"id" gorm:"column:id;"`
	Description string `json:"description" gorm:"column:description;"`
	GroupName   string `json:"groupName" gorm:"column:groupName"`
}

func (*Feature) TableName() string {
	return common.TableFeature
}

var (
	ErrFeatureViewNoPermission = common.ErrNoPermission(
		errors.New("Bạn không có quyền xem các chức năng"),
	)
)
