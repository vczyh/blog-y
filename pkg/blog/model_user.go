package blog

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}
