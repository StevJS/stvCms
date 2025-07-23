package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	Author  string
	//Images  []byte `gorm:"type:bytea"`
	CodeContent []CodeContent
}

type CodeContent struct {
	gorm.Model
	Code     string
	Language string

	//relaciones
	PostID uint
	Post   Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
