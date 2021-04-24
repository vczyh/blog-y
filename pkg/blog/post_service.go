package blog

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type PostReq struct {
	Title     *string `json:"title"`
	Subtitle  *string `json:"subtitle"`
	Content   *string `json:"content"`
	CoverURL  *string `json:"cover_url"`
	CoverDesc *string `json:"cover_desc"`
}

type PostRes struct {
	PostId    string    `json:"post_id"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Content   string    `json:"content"`
	CoverURL  string    `json:"cover_url"`
	CoverDesc string    `json:"cover_desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      []TagRes  `json:"tags"`
}

// CreatePostService 创建文章
func CreatePostService(req PostReq) (string, error) {
	var p Post
	if req.Title != nil {
		p.Title = *req.Title
	}
	if req.Subtitle != nil {
		p.Subtitle = *req.Subtitle
	}
	if req.Content != nil {
		p.Content = *req.Content
	}
	if req.CoverURL != nil {
		p.CoverURL = *req.CoverURL
	}
	if req.CoverDesc != nil {
		p.CoverDesc = *req.CoverDesc
	}

	postId, err := NewPost(p)
	if err != nil {
		l.Error("create post to mysql error", "error", err)
		return "", err
	}

	return postId, nil
}

// PostListService 分页文章列表
func PostListService(page, pageSize int, tagId string) (*PageInfo, error) {
	var ps []Post
	var pageInfo *PageInfo

	err := db.Transaction(func(tx *gorm.DB) error {

		var postIds []string
		var flag bool

		// tag condition
		if tagId != "" {
			flag = true
			rows, err := tx.Table("post_tag").Where("tag_id = ?", tagId).Select("post_id").Rows()
			if err != nil {
				return err
			}
			defer func() { _ = rows.Close() }()
			for rows.Next() {
				var postId string
				if err = rows.Scan(&postId); err != nil {
					return err
				}
				postIds = append(postIds, postId)
			}
		}

		// posts
		var d = tx
		if flag {
			d = d.Where("post_id in (?)", postIds)
		}
		d = d.Preload("Tags").Order("created_at DESC")
		p, err := Page(d, &ps, page, pageSize)
		if err != nil {
			return err
		}

		pageInfo = p

		return nil
	})
	if err != nil {
		return nil, err
	}

	var postsRes []PostRes
	for _, p := range ps {
		var tagsRes []TagRes
		for _, t := range p.Tags {
			tagsRes = append(tagsRes, TagRes{
				TagId: t.TagId,
				Name:  t.Name,
			})
		}
		postsRes = append(postsRes, PostRes{
			PostId:    p.PostId,
			Title:     p.Title,
			Subtitle:  p.Subtitle,
			Content:   p.Content,
			CoverURL:  p.CoverURL,
			CoverDesc: p.CoverDesc,
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
			Tags:      tagsRes,
		})
	}
	pageInfo.List = postsRes

	return pageInfo, nil
}

// GetPostByPostIdService 根据postId查询文章
func GetPostByPostIdService(id string) (PostRes, error) {
	// post
	var postRes PostRes
	p, err := GetPostByPostId(id)
	if err != nil {
		l.Error("failed to get post by post_id", "post_id", id)
		return postRes, err
	}

	var tagsRes []TagRes
	for _, t := range p.Tags {
		tagsRes = append(tagsRes, TagRes{
			TagId: t.TagId,
			Name:  t.Name,
		})
	}

	return PostRes{
		PostId:    p.PostId,
		Title:     p.Title,
		Subtitle:  p.Subtitle,
		Content:   p.Content,
		CoverURL:  p.CoverURL,
		CoverDesc: p.CoverDesc,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Tags:      tagsRes,
	}, nil
}

// UpdatePostService 更新文章
func UpdatePostService(postId string, req PostReq) error {
	if postId == "" {
		return fmt.Errorf("param empty")
	}

	// 需要更新的字段
	var m = make(map[string]interface{}, 5)
	if req.Title != nil {
		m["title"] = *req.Title

	}
	if req.Subtitle != nil {
		m["subtitle"] = *req.Subtitle
	}
	if req.Content != nil {
		m["content"] = *req.Content
	}
	if req.CoverURL != nil {
		m["cover_url"] = *req.CoverURL
	}
	if req.CoverDesc != nil {
		m["cover_desc"] = *req.CoverDesc
	}

	tx := db.Model(&Post{}).Where("post_id = ?", postId).Updates(m)
	if err := tx.Error; err != nil {
		l.Error("failed to update post", "post_id", postId, "error", err)
		return err
	}

	return nil
}

// AddPostTagService 添加post的tag
func AddPostTagService(postId string, ids []string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if id == "" {
				return fmt.Errorf("tagId must be not empty")
			}
			// tag 是否存在
			var t Tag
			res := tx.Debug().Where("tag_id = ?", id).Find(&t)
			if err := res.Error; err != nil {
				return err
			}
			if res.RowsAffected == 0 {
				continue
			}
			// 关联记录是否存在
			rows, err := tx.Debug().Table("post_tag").Where("post_id = ? AND tag_id = ?", postId, id).Rows()
			if err != nil {
				return err
			}
			defer func() { _ = rows.Close() }()
			for rows.Next() {
				continue
			}
			// 添加关联
			if err := tx.Debug().Table("post_tag").Create(map[string]interface{}{
				"post_id": postId,
				"tag_id":  id,
			}).Error; err != nil {
				return err
			}

			//if err := tx.Model(&Post{PostId: postId}).Where("post_id = ?", postId).
			//	Association("Tags").Append(&Tag{TagId: t.TagId}); err != nil {
			//	return err
			//}
		}

		return nil
	})

	if err != nil {
		l.Error("failed to add post tag", "postId", postId, "ids", ids, "error", err)
		return err
	}

	return nil
}

// DeletePostTagService 删除post的tag
func DeletePostTagService(postId string, tagIds []string) error {
	var tags []*Tag
	for _, id := range tagIds {
		tags = append(tags, &Tag{TagId: id})

	}

	err := db.Model(&Post{PostId: postId}).Association("Tags").Delete(tags)
	if err != nil {
		l.Error("failed to delete post's tag", "postId", postId, "tagIds", tagIds)
		return err
	}

	return nil
}

// TimelineService Timeline 时间轴
func TimelineService() ([]interface{}, error) {
	var timeline []interface{}

	err := db.Transaction(func(tx *gorm.DB) error {

		// year DESC
		var years []int
		rows, err := tx.Raw("SELECT YEAR(created_at) AS year FROM post GROUP BY year ORDER BY year DESC").Rows()
		if err != nil {
			return err
		}
		defer func() { _ = rows.Close() }()
		for rows.Next() {
			var y int
			if err := rows.Scan(&y); err != nil {
				return err
			}
			years = append(years, y)
		}

		// posts
		for _, y := range years {
			var ps []Post
			res := tx.Where("YEAR(created_at) = ?", y).Order("created_at DESC").Preload("Tags").Find(&ps)
			if err := res.Error; err != nil {
				return err
			}
			var postsRes []PostRes
			for _, p := range ps {
				var tagsRes []TagRes
				for _, t := range p.Tags {
					tagsRes = append(tagsRes, TagRes{
						TagId: t.TagId,
						Name:  t.Name,
					})
				}
				postsRes = append(postsRes, PostRes{
					PostId:    p.PostId,
					Title:     p.Title,
					Subtitle:  p.Subtitle,
					Content:   p.Content,
					CoverURL:  p.CoverURL,
					CoverDesc: p.CoverDesc,
					CreatedAt: p.CreatedAt,
					UpdatedAt: p.UpdatedAt,
					Tags:      tagsRes,
				})
			}
			timeline = append(timeline, map[string]interface{}{
				"year": y,
				"list": postsRes,
			})
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return timeline, nil
}

// DeletePostService 删除文章
func DeletePostService(postId string) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Exec("DELETE FROM post_tag WHERE post_id = ?", postId).Error; err != nil {
			return err
		}

		if err := tx.Debug().Where("post_id = ?", postId).Delete(&Post{}).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}
