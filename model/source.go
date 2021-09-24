package model

type SourceModel struct {
	BaseModel
	Title  string `json:"title" gorm:"column:title"`
	Domain string `json:"domain" gorm:"column:domain"`
}

func (u *SourceModel) TableName() string {
	return "z_source"
}

func (u *SourceModel) Create() error {
	return DB.Self.Create(&u).Error
}
