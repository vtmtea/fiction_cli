package model

type SourceRuleModel struct {
	SourceID      int    `json:"source_id" gorm:"column:source_id"`
	ListLink      string `json:"list_link" gorm:"column:list_link"`
	BookListRoute string `json:"book_list_route" gorm:"column:book_list_route"`
	BookLinkRoute string `json:"book_link_route" gorm:"column:book_link_route"`
	BookNameRoute string `json:"book_name_route" gorm:"column:book_name_route"`
	AuthorRoute   string `json:"author_route" gorm:"column:author_route"`
}

func (u *SourceRuleModel) TableName() string {
	return "z_source_rule"
}

func (u *SourceRuleModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetSourceRule(sourceId uint64) ([]SourceRuleModel, error) {
	var u []SourceRuleModel
	err := DB.Self.Where("source_id = ?", sourceId).Find(&u).Error
	return u, err
}
