package blog

import "gorm.io/gorm"

type Config struct {
	gorm.Model

	Name         string
	InitialValue string
	CurrentValue string
}

func (Config) TableName() string {
	return "config"
}
