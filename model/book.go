package model

import "time"

type BookModel struct {
	BaseModel
	Title           string    `json:"title" gorm:"column:title"`
	AuthorID        int       `json:"author_id" gorm:"column:author_id"`
	CategoryID      int       `json:"category_id" gorm:"column:category_id"`
	CoverImage      string    `json:"cover_image" gorm:"column:cover_image"`
	Description     string    `json:"description" gorm:"column:description"`
	Rate            int       `json:"rate" gorm:"column:rate"`
	RatePeopleCount int       `json:"rate_people_count" gorm:"column:rate_people_count"`
	ClickCount      int       `json:"click_count" gorm:"column:click_count"`
	CollectCount    int       `json:"collect_count" gorm:"column:collect_count"`
	RecommendCount  int       `json:"recommend_count" gorm:"column:recommend_count"`
	LastUpdateTime  time.Time `json:"last_update_time" gorm:"column:last_update_time"`
	BookStatus      int       `json:"book_status" gorm:"column:book_status"`
}

func (u *BookModel) TableName() string {
	return "z_book"
}

func (u *BookModel) Create() error {
	return DB.Self.Create(&u).Error
}
