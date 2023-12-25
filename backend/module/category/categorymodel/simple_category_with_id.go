package categorymodel

type SimpleCategoryWithId struct {
	CategoryId string `json:"categoryId" gorm:"column:categoryId;"`
}
