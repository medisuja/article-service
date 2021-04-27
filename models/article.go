package models

import "time"

type (
	Article struct {
		ID        int       `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key,column:id"`
		Author    string    `json:"author" gorm:"column:author"`
		Title     string    `json:"title" gorm:"column:title"`
		Slug      string    `json:"slug" gorm:"column:slug"`
		Body      string    `json:"body" gorm:"column:body"`
		CreatedAt time.Time `json:"created" gorm:"column:created" sql:"DEFAULT:current_timestamp"`
	}

	AlgoliaArticle struct {
		ObjectID  int    `json:"objectID"`
		Author    string `json:"author"`
		Title     string `json:"title"`
		Slug      string `json:"slug"`
		Body      string `json:"body"`
		CreatedAt int64  `json:"created"`
	}
)

func (Article) TableName() string {
	return "Article"
}
