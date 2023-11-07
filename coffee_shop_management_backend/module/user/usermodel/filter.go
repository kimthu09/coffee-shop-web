package usermodel

type Filter struct {
	SearchKey string `json:"searchKey,omitempty" form:"search"`
	IsActive  *bool  `json:"active,omitempty" form:"active"`
	Role      string `json:"role,omitempty" form:"role"`
}
