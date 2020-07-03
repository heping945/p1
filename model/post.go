package model

import (
	"github.com/jinzhu/gorm"
)

type (
	Topic struct {
		gorm.Model
		Title      string   `json:"title" gorm:"not null;unique" binding:"required"`
		Difficulty uint8    `json:"difficulty" gorm:"not null" binding:"required,max=3,min=1"`
		Slug       string   `json:"slug" gorm:"type:varchar(140);unique_index;not null"`
		Md         string   `json:"md" binding:"required" gorm:"size:2048"`
		Tags       []Tag    `gorm:"many2many:topic_tags;" json:"tags"`
		Post       []Post   `json:"post" `
		SourceUrl  string   `json:"source_url" binding:"url"`
		Similar    []*Topic `gorm:"many2many:similar_topics;association_jointable_foreignkey:similar_id" json:"similar"`
	}

	TransFormTopic struct {
		ID         uint   `json:"id"`
		Title      string `json:"title"`
		Difficulty uint8  `json:"difficulty"`
		Slug       string `json:"slug"`
	}
)

type (
	Post struct {
		gorm.Model
		Category string `json:"category" binding:"required" gorm:"unique_index:ct"`
		Md       string `json:"md" binding:"required" gorm:"size:4096"`
		TopicID  uint   `json:"topic_id" binding:"required" gorm:"unique_index:ct"`
	}
	TransFormPost struct {
		ID           uint   `json:"id"`
		Category     string `json:"category"`
		TopicID      uint   `json:"topic_i"`
		ShortMd      string `json:"short_md"`
		TopicIfExist bool   `json:"topic_if_exist"`
	}
)

type Tag struct {
	gorm.Model
	Name string `json:"name" gorm:"not null;type:varchar(40)"  binding:"required"`
}

type Category struct {
	ID    uint   `json:"id"`
	Label string `json:"label" binding:"required" `
	Name  string `json:"name" binding:"required" gorm:"not null;unique"`
}
