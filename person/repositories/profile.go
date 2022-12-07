package repositories

import (
	"party/database/models"

	"gorm.io/gorm"
)

type ProfileRepositories interface {
	GetProfile(id int) (models.Profile, error)
}

func RepositoriesProfile(db *gorm.DB) *repositories{
	return &repositories{db}
}

func (r *repositories) GetProfile(id int) (models.Profile, error) {
	var profile models.Profile
     err := r.DB.First(&profile, id).Error

    return profile, err
}