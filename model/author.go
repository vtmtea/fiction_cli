package model

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
