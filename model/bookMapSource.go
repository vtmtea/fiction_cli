package model

type BookMapSourceModel struct {
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	BookID     uint64 `json:"book_id" gorm:"column:book_id"`
	SourceID   uint64 `json:"source_id" gorm:"column:source_id"`
	SourceLink string `json:"source_link" gorm:"column:source_link"`
}

func (u *BookMapSourceModel) TableName() string {
	return "z_book_map_source"
}

func (u *BookMapSourceModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetSingleData(bookId uint64, sourceId uint64) (BookMapSourceModel, error) {
	u := BookMapSourceModel{BookID: bookId, SourceID: sourceId}
	err := DB.Self.Where("book_id = ? and source_id = ?", bookId, sourceId).First(&u).Error
	return u, err
}
