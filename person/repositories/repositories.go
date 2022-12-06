package repositories

import "gorm.io/gorm"

type repositories struct {
	DB *gorm.DB
}