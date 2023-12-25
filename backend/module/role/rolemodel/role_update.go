package rolemodel

import "coffee_shop_management_backend/common"

type RoleUpdate struct {
	Name     *string   `json:"name" gorm:"column:name;"`
	Features *[]string `json:"features" gorm:"-"`
}

func (*RoleUpdate) TableName() string {
	return common.TableRole
}

func (data *RoleUpdate) Validate() *common.AppError {
	if data.Name != nil && common.ValidateEmptyString(*data.Name) {
		return ErrRoleNameEmpty
	}
	if data.Features != nil {
		if len(*data.Features) == 0 {
			return ErrRoleFeaturesEmpty
		}
		for _, v := range *data.Features {
			if !common.ValidateNotNilId(&v) {
				return ErrRoleFeatureInvalid
			}
		}
	}
	return nil
}
