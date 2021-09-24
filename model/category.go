package model

type CategoryModel struct {
	BaseModel
	Title     string `json:"title" gorm:"column:title"`
	Order     int    `json:"order" gorm:"column:order"`
	TextMap   string `json:"text_map" gorm:"column:text_map"`
	BookCount int    `json:"book_count" gorm:"column:book_count"`
}

func (u *CategoryModel) TableName() string {
	return "z_category"
}

func (u *CategoryModel) Create() error {
	return DB.Self.Create(&u).Error
}
