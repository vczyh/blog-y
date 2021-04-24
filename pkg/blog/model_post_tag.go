package blog

//type PostTag struct {
//	gorm.Model
//
//	PostId string
//	TagId  string
//}
//
//func (PostTag) TableName() string {
//	return "post_tag"
//}
//
//func GetTagsByPostId(postId string) ([]Tag, error) {
//	var ts []Tag
//	tagIdsQuery := db.Model(&PostTag{}).Select("tag_id").Where("post_id = ?", postId)
//	return ts, db.Debug().Where("tag_id in (?)", tagIdsQuery).Find(&ts).Error
//}

//func DeleteMappingByPostId(postId string,tagIds []string) {
//	tagIdsQuery := db.Model(&PostTag{}).Select("tag_id").Where("post_id = ?", postId)
//	db.Debug().Where("tag_id in (?)",tagIds).Delete()
//}
