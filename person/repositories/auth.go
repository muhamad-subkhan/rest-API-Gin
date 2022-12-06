package repositories

import (
	"party/database/models"

	"gorm.io/gorm"
)

type AuthRepositories interface {
	Register(profile models.Profile) (models.Profile, error)
	Login(email string) (models.Profile, error)
}

func RepositoriesAuth(db *gorm.DB) *repositories {
	return &repositories{db}
}

func (r *repositories) Register(profile models.Profile) (models.Profile, error) {
	 err := r.DB.Create(&profile).Error 
     return profile, err
}

func (r *repositories) Login(email string) (models.Profile, error) {
	var profile models.Profile
    err := r.DB.Where("email =?", email).First(&profile).Error
    return profile, err
}