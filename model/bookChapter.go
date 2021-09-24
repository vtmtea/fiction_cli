package model

type BookChapterModel struct {
	BaseModel
	Title      string `json:"title" gorm:"column:title"`
	SourceID   int    `json:"source_id" gorm:"column:source_id"`
	BookID     int    `json:"book_id" gorm:"column:book_id"`
	SourceLink string `json:"source_link" gorm:"column:source_link"`
	Content    string `json:"content" gorm:"column:content"`
}

func (u *BookChapterModel) TableName() string {
	return "z_book_chapter"
}

func (u *BookChapterModel) Create() error {
	return DB.Self.Create(&u).Error
}
