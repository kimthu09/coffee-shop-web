package ingredientmodel

type Filter struct {
	SearchKey   string   `json:"searchKey,omitempty" form:"search"`
	MinPrice    *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice    *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	MinAmount   *float32 `json:"minAmount,omitempty" form:"minAmount"`
	MaxAmount   *float32 `json:"maxAmount,omitempty" form:"maxAmount"`
	MeasureType string   `json:"measureType,omitempty" form:"measureType"`
}
