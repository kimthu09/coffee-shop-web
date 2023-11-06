package productmodel

import (
	"coffee_shop_management_backend/common"
	"coffee_shop_management_backend/module/sizefood/sizefoodmodel"
)

type FoodUpdate struct {
	*ProductUpdate `json:",inline"`
	Categories     *[]string                       `json:"categories" gorm:"-"`
	Sizes          *[]sizefoodmodel.SizeFoodUpdate `json:"sizes" gorm:"-"`
}

func (*FoodUpdate) TableName() string {
	return common.TableFood
}

func (data *FoodUpdate) Validate() error {
	if err := (*data.ProductUpdate).Validate(); err != nil {
		return err
	}
	if data.Categories != nil {
		if len(*data.Categories) == 0 {
			return ErrFoodCategoryEmpty
		}
		mapExistCategory := make(map[string]int)
		for _, v := range *data.Categories {
			mapExistCategory[v]++
			if mapExistCategory[v] == 2 {
				return ErrFoodExistDuplicateCategory
			}
		}
	}
	if data.Sizes != nil {
		if len(*data.Sizes) == 0 {
			return ErrFoodSizeEmpty
		}
		mapExistSize := make(map[string]int)
		for _, v := range *data.Sizes {
			if err := v.Validate(); err != nil {
				return err
			}
			if v.SizeId != nil {
				mapExistSize[*v.SizeId]++
				if mapExistSize[*v.SizeId] == 2 {
					return ErrFoodExistDuplicateSize
				}
			}
		}
	}
	return nil
}
