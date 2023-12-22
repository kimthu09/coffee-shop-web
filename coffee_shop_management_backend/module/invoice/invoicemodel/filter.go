package invoicemodel

type Filter struct {
	SearchKey         string  `json:"searchKey,omitempty" form:"search"`
	DateFromCreatedAt *int64  `json:"createdAtFrom,omitempty" form:"createdAtFrom"`
	DateToCreatedAt   *int64  `json:"createdAtTo,omitempty" form:"createdAtTo"`
	MinPrice          *int    `json:"minPrice,omitempty" form:"minPrice"`
	MaxPrice          *int    `json:"maxPrice,omitempty" form:"maxPrice"`
	CreatedBy         *string `json:"createdBy,omitempty" form:"createdBy"`
	Customer          *string `json:"customer,omitempty" form:"customer"`
}
