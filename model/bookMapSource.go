package model

type BookMapSourceModel struct {
	ID         uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	BookID     int    `json:"book_id" gorm:"column:book_id"`
	SourceID   int    `json:"source_id" gorm:"column:source_id"`
	SourceLink string `json:"source_link" gorm:"column:source_link"`
}

func (u *BookMapSourceModel) TableName() string {
	return "z_book_map_source"
}

func (u *BookMapSourceModel) Create() error {
	return DB.Self.Create(&u).Error
}
