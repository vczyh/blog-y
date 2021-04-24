package blog

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	gorm.Model

	TagId string `gorm:"size:255;index:,unique"`
	Name  string `gorm:"index:,unique"`
	//Posts []Post
}

func (Tag) TableName() string {
	return "tag"
}

func NewTag(name string) (string, error) {
	var tag = Tag{
		TagId: fmt.Sprintf("%d", time.Now().Unix()),
		Name:  name,
	}
	return tag.TagId, db.Create(&tag).Error
}
