package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	BookID  uint   `json:"book_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
