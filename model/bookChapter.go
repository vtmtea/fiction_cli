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

func (u *BookChapterModel) SaveContent() {
	DB.Self.Model(&u).Update("content", u.Content)
}

func BatchCreate(models []BookChapterModel) error {
	return DB.Self.Create(&models).Error
}

func GetChapterById(id uint64) BookChapterModel {
	var u BookChapterModel
	DB.Self.First(&u, id)
	return u
}

func GetChapterCount(sourceId uint64, bookId uint64) int {
	var count int
	err := DB.Self.Model(&BookChapterModel{}).Where("book_id = ? and source_id = ?", bookId, sourceId).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}
