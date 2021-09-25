package model

type BookChapterModel struct {
	BaseModel
	Title      string `json:"title" gorm:"column:title"`
	SourceID   uint64 `json:"source_id" gorm:"column:source_id"`
	BookID     uint64 `json:"book_id" gorm:"column:book_id"`
	SourceLink string `json:"source_link" gorm:"column:source_link"`
	Content    string `json:"content" gorm:"column:content"`
}

func (u *BookChapterModel) TableName() string {
	return "z_book_chapter"
}

func (u *BookChapterModel) Create() error {
	return DB.Self.Create(&u).Error
}

func BatchCreate(models []BookChapterModel) error {
	return DB.Self.Create(&models).Error
}
