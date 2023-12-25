package productmodel

type Filter struct {
	SearchKey string `json:"searchKey,omitempty" form:"search"`
	IsActive  *bool  `json:"active" form:"active"`
}
