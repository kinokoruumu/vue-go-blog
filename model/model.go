package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Text  string `json:"text"`
}
