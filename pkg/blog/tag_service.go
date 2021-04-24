package blog

import "fmt"

type TagReq struct {
	Name string `json:"name"`
}

type TagRes struct {
	TagId string `json:"tag_id"`
	Name  string `json:"name"`
}

func CreateTagService(req TagReq) (string, error) {
	if req.Name == "" {
		return "", fmt.Errorf("tag name must be not empty")
	}
	var tag Tag
	res := db.Where("name = ?", req.Name).Find(&tag)
	if err := res.Error; err != nil {
		l.Error("failed to get tag by name", "error", err)
		return "", err
	}
	if res.RowsAffected == 1 {
		return tag.TagId, nil
	}

	tagId, err := NewTag(req.Name)
	if err != nil {
		l.Error("failed to create tag", "error", err)
		return "", err
	}
	return tagId, nil
}

func TagListService() ([]TagRes, error) {
	var ts []Tag
	if err := db.Find(&ts).Error; err != nil {
		l.Error("failed to get tags", "error", err)
		return nil, err
	}

	var res []TagRes
	for _, t := range ts {
		res = append(res, TagRes{
			TagId: t.TagId,
			Name:  t.Name,
		})
	}

	return res, nil
}
