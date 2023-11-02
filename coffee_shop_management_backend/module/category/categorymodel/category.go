package categorymodel

import (
	"coffee_shop_management_backend/common"
	"errors"
	"time"
)

type SimpleCategory struct {
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}

type Category struct {
	Id            string     `json:"id" gorm:"column:id;"`
	Name          string     `json:"name" gorm:"column:name;"`
	Description   string     `json:"description" gorm:"column:description;"`
	CreatedAt     *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt;"`
	AmountProduct int        `json:"amountProduct" gorm:"column:amountProduct;"`
}

func (*Category) TableName() string {
	return common.TableCategory
}

var (
	ErrIdInvalid = common.NewCustomError(
		errors.New("id of category is invalid"),
		"id of category is invalid",
		"ErrIdInvalid",
	)
	ErrNameEmpty = common.NewCustomError(
		errors.New("name of category is empty"),
		"name of category is empty",
		"ErrNameEmpty",
	)
	ErrAmountProductCategoryNotExist = common.NewCustomError(
		errors.New("amount product of category is empty"),
		"amount product of category is empty",
		"ErrAmountProductCategoryNotExist",
	)
)
