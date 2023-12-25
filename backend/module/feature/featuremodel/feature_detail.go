package featuremodel

type FeatureDetail struct {
	Id          string `json:"id" gorm:"column:id;"`
	Description string `json:"description" gorm:"column:description;"`
	GroupName   string `json:"groupName" gorm:"column:groupName"`
	IsHas       bool   `json:"isHas" gorm:"-"`
}
