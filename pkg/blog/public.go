package blog

import (
	"blog-y/pkg/common/config"
	"blog-y/pkg/common/log"
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

// log
var l *log.Logger

func WithLogger(logger *log.Logger) {
	l = logger
}

// config
var c *config.Config

func WithConfig(config *config.Config) {
	c = config
}

// MySQL
var db *gorm.DB

func WithMySQL(gormDB *gorm.DB) {
	db = gormDB
	//db.AutoMigrate(&User{}, &Tag{}, &Post{}, &Config{})
}

type PageInfo struct {
	Total      int64       `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	PageNumber int64       `json:"page_number"`
	List       interface{} `json:"list"`
}

func Page(db *gorm.DB, v interface{}, page, pageSize int) (*PageInfo, error) {
	t := reflect.TypeOf(v)

	if t.Kind() != reflect.Ptr {
		return nil, fmt.Errorf("v must be ptr type")
	}

	if t = t.Elem(); t.Kind() != reflect.Slice {
		return nil, fmt.Errorf("*v must be slice type")
	}

	model := reflect.New(t.Elem()).Interface()

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var pageInfo = PageInfo{
		Page:     page,
		PageSize: pageSize,
	}

	// 执行两条SQL
	if err := db.Debug().Model(model).Count(&pageInfo.Total).Limit(pageSize).Offset((page - 1) * pageSize).Find(v).Error
		err != nil {
		return nil, err
	}

	pageInfo.PageNumber = (pageInfo.Total-1)/int64(pageInfo.PageSize) + 1
	pageInfo.List = v

	return &pageInfo, nil
}
