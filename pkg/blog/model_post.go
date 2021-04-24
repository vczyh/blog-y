package blog

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model

	PostId    string `gorm:"size:255;index:,unique"`
	Title     string
	Subtitle  string
	Content   string
	CoverURL  string `gorm:"column:cover_url"`
	CoverDesc string
	Tags      []Tag `gorm:"many2many:post_tag;foreignKey:PostId;joinForeignKey:PostId;References:TagId;JoinReferences:TagId"`
}

func (Post) TableName() string {
	return "post"
}

func NewPost(post Post) (string, error) {
	post.PostId = fmt.Sprintf("%d", time.Now().Unix())
	return post.PostId, db.Create(&post).Error
}

func GetPostByPostId(id string) (Post, error) {
	var p Post
	return p, db.Debug().Where("post_id = ?", id).Preload("Tags").First(&p).Error
}
