package rolefeaturemodel

import (
	"coffee_shop_management_backend/common"
)

type RoleFeature struct {
	RoleId    string `json:"roleId" gorm:"column:roleId;"`
	FeatureId string `json:"featureId" gorm:"column:featureId;"`
}

func (*RoleFeature) TableName() string {
	return common.TableRoleFeature
}
