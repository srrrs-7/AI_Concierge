package model

type Author struct {
	ID     int64  `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Number int64  `gorm:"column:number" json:"number"`
	Bio    string `gorm:"column:bio" json:"bio"`
}
