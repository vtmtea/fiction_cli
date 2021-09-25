package model

import "github.com/jinzhu/gorm"

type AuthorModel struct {
	ID        int    `json:"-" gorm:"primary_key;AUTO_INCREMENT;column:id"`
	Name      string `json:"name" gorm:"column:name"`
	BookCount int    `json:"book_count" gorm:"column:book_count"`
}

func (u *AuthorModel) TableName() string {
	return "z_author"
}

func (u *AuthorModel) Create() error {
	return DB.Self.Create(&u).Error
}

func (u *AuthorModel) UpdateCount() error {
	return DB.Self.Model(&u).UpdateColumn("book_count", gorm.Expr("book_count + ?", 1)).Error
}

func GetAuthor(name string) (AuthorModel, error) {
	u := AuthorModel{Name: name}
	err := DB.Self.Where(&u).First(&u).Error
	return u, err
}
