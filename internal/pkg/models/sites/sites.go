package models

type WebSite struct {
	Name string `gorm:"column:name;not null;size:50" json:"name"`
	Link string `gorm:"column:link;not null;size:255" json:"link"`
	Code int    `gorm:"column:code;not null;default:200" json:"code"`
}
