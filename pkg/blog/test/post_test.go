package test

import (
	"blog-y/pkg/blog"
	"blog-y/pkg/common/mysql"
	"fmt"
	"gorm.io/gorm"
	"testing"
	"time"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	var err error
	db, err = mysql.New(
		"10.0.44.13",
		3306,
		"root",
		"Unicloud1221@",
		"blog")
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestCreatePost(t *testing.T) {
	tx := db.Create(&blog.Post{
		PostId: fmt.Sprintf("test-%d", time.Now().Unix()),
		Tags: []blog.Tag{
			{TagId: fmt.Sprintf("test-%d", time.Now().Unix()), Name: "test-tag2"},
		},
	})
	if err := tx.Error; err != nil {
		t.Fatal(err)
	}
}

func TestGetPostByPostId(t *testing.T) {
	postId := "test-1618232636"
	var p blog.Post
	res := db.Debug().Where("post_id = ?", postId).Preload("Tags").Find(&p)
	if err := res.Error; err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.RowsAffected)
	fmt.Println(p)
}

func TestGetPosts(t *testing.T) {
	var ps []blog.Post
	tx := db.Preload("Tags").Find(&ps)
	if err := tx.Error; err != nil {
		t.Fatal(err)
	}
	for _, p := range ps {
		fmt.Println(p.Tags)
	}
}

// 只删除m2m关联记录 更新post的updatedAt
func TestDeleteTags(t *testing.T) {
	tagId1 := "test-1618232636"
	tagId2 := "test-1618233125"
	postId := "test-1618233125"
	err := db.Debug().Model(&blog.Post{PostId: postId}).
		Association("Tags").
		Delete(&blog.Tag{TagId: tagId1}, &blog.Tag{TagId: tagId2})
	if err != nil {
		t.Fatal(err)
	}
}

// 添加m2m关联 不更新post的updatedAt
func TestAddTags(t *testing.T) {
	tagId1 := "test-1618232636"
	tagId2 := "test-1618233125"
	postId := "test-1618233125"

	d := db.Debug().Model(&blog.Post{PostId: postId})

	d = d.Where("post_id = ?", postId)

	err := d.Association("Tags").Append(&blog.Tag{TagId: tagId1},
		&blog.Tag{TagId: tagId2},
		&blog.Tag{TagId: fmt.Sprintf("test-%d", time.Now().Unix()), Name: "test-tag3"})

	if err != nil {
		t.Fatal(err)
	}
}

// 晴空关联
