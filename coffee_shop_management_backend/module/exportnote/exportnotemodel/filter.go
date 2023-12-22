package exportnotemodel

type Filter struct {
	SearchKey         string        `json:"searchKey,omitempty" form:"search"`
	DateFromCreatedAt *int64        `json:"createdAtFrom,omitempty" form:"createdAtFrom"`
	DateToCreatedAt   *int64        `json:"createdAtTo,omitempty" form:"createdAtTo"`
	CreatedBy         *string       `json:"createdBy,omitempty" form:"createdBy"`
	Reason            *ExportReason `json:"reason,omitempty" form:"reason"`
}
