package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
)

type FoodCreate struct {
	*ProductCreate `json:",inline"`
	Categories     []string                       `json:"categories" gorm:"-"`
	Sizes          []sizefoodmodel.SizeFoodCreate `json:"sizes" gorm:"-"`
}

func (*FoodCreate) TableName() string {
	return common.TableFood
}

func (data *FoodCreate) Validate() error {
	if err := (*data.ProductCreate).Validate(); err != nil {
		return err
	}
	if data.Categories == nil || len(data.Categories) == 0 {
		return ErrCategoryEmpty
	}
	mapExistCategory := make(map[string]int)
	for _, v := range data.Categories {
		mapExistCategory[v]++
		if mapExistCategory[v] == 2 {
			return ErrExistDuplicateCategory
		}
	}
	if data.Sizes == nil || len(data.Sizes) == 0 {
		return ErrSizeEmpty
	}
	for _, v := range data.Sizes {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
