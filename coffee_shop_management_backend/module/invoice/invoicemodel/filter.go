package invoicemodel

type Filter struct {
	SearchKey string   `json:"searchKey,omitempty" form:"search"`
	MinPrice  *float32 `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice  *float32 `json:"maxPrice,omitempty" form:"maxPrice"`
	IsHasDebt *bool    `json:"hasDebt" form:"hasDebt"`
}
