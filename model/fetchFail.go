package model

type FetchFailModel struct {
	BaseModel
	SourceId   uint64 `json:"source_id" gorm:"column:source_id"`
	SourceLink string `json:"source_link" gorm:"column:source_link"`
}

func (u *FetchFailModel) TableName() string {
	return "z_fetch_fail"
}

func (u *FetchFailModel) Create() error {
	return DB.Self.Create(&u).Error
}

func GetFailList(count int) []FetchFailModel {
	var fetchFails []FetchFailModel
	DB.Self.Limit(count).Find(&fetchFails)
	return fetchFails
}

func DeleteRecord(id uint64) {
	DB.Self.Delete(&FetchFailModel{}, id)
}
