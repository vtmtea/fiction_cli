package model

type SourceModel struct {
	BaseModel
	Title               string `json:"title" gorm:"column:title"`
	Domain              string `json:"domain" gorm:"column:domain"`
	BookNameRoute       string `json:"book_name_route" gorm:"column:book_name_route"`
	AuthorRoute         string `json:"author_route" gorm:"column:author_route"`
	CategoryRoute       string `json:"category_route" gorm:"column:category_route"`
	DescriptionRoute    string `json:"description_route" gorm:"column:description_route"`
	LastUdpateRoute     string `json:"last_udpate_route" gorm:"column:last_udpate_route"`
	CoverRoute          string `json:"cover_route" gorm:"column:cover_route"`
	StatusRoute         string `json:"status_route" gorm:"column:status_route"`
	ChapterListRoute    string `json:"chapter_list_route" gorm:"column:chapter_list_route"`
	ChapterContentRoute string `json:"chapter_content_route" gorm:"column:chapter_content_route"`
}

func (u *SourceModel) TableName() string {
	return "z_source"
}

func (u *SourceModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetSourceById(id uint64) (SourceModel, error) {
	var u SourceModel
	err := DB.Self.First(&u, id).Error
	return u, err
}
