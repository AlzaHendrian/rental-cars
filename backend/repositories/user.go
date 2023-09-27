package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := repo.db.Find(&users).Error

	return users, err
}

func (repo *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := repo.db.First(&user, ID).Error

	return user, err
}

func (repo *repository) CreateUser(user models.User) (models.User, error) {
	err := repo.db.Create(&user).Error

	return user, err
}

func (repo *repository) UpdateUser(user models.User) (models.User, error) {
	err := repo.db.Save(&user).Error

	return user, err
}

func (repo *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := repo.db.Delete(&user, ID).Scan(&user).Error

	return user, err
}
