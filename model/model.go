package model

import "time"

type BaseModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"-"`
	Status    uint64    `gorm:"column:status" json:"-"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"-"`
}

type Token struct {
	Token string `json:"token"`
}
