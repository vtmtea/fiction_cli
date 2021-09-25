package model

import "github.com/jinzhu/gorm"

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

func (u *CategoryModel) UpdateCount() error {
	return DB.Self.Model(&u).UpdateColumn("book_count", gorm.Expr("book_count + ?", 1)).Error
}

func GetCategoryByMap(text string) (CategoryModel, error) {
	u := CategoryModel{}
	err := DB.Self.Where("text_map like ?", "%"+text+"%").First(&u).Error
	return u, err
}
