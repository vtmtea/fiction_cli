package model

import "github.com/jinzhu/gorm"

type BookModel struct {
	BaseModel
	Title           string `json:"title" gorm:"column:title"`
	AuthorID        int    `json:"author_id" gorm:"column:author_id"`
	CategoryID      uint64 `json:"category_id" gorm:"column:category_id"`
	CoverImage      string `json:"cover_image" gorm:"column:cover_image"`
	Description     string `json:"description" gorm:"column:description"`
	Rate            int    `json:"rate" gorm:"column:rate"`
	RatePeopleCount int    `json:"rate_people_count" gorm:"column:rate_people_count"`
	ClickCount      int    `json:"click_count" gorm:"column:click_count"`
	CollectCount    int    `json:"collect_count" gorm:"column:collect_count"`
	RecommendCount  int    `json:"recommend_count" gorm:"column:recommend_count"`
	LastUpdateTime  string `json:"last_update_time" gorm:"column:last_update_time"`
	BookStatus      int    `json:"book_status" gorm:"column:book_status"`
}

func (u *BookModel) TableName() string {
	return "z_book"
}

func (u *BookModel) Create() error {
	return DB.Self.Create(&u).Error
}

func BookExist(sourceLink string) bool {
	var bookMapSource BookMapSourceModel
	err := DB.Self.Where("source_link = ?", sourceLink).First(&bookMapSource).Error
	if err != nil && err == gorm.ErrRecordNotFound {
		return false
	}
	if bookMapSource.ID > 0 {
		return true
	}
	return false
}

func GetBookById(id uint64) BookModel {
	var u BookModel
	DB.Self.First(&u, id)
	return u
}

func GetBookByMapAttr(attr map[string]interface{}) (BookModel, error) {
	var u BookModel
	err := DB.Self.Where(attr).First(&u).Error
	return u, err
}
