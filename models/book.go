package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ID     int    `json: "id" form:"id"`
	Title  string `json: "title" form:"title"`
	ISBN   string `json: "isbn" form:"isbn"`
	Writer string `json: "writer" form:"writer"`
}
