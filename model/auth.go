package model

import "github.com/jinzhu/gorm"

type AuthenCode struct {
	gorm.Model
	Code string `gorm:"not null;unique_index" json:"code" binding:"required"`
}
