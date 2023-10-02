package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	FindCars() ([]models.Car, error)
	GetCar(ID int) (models.Car, error)
	CreateCar(car models.Car) (models.Car, error)
	UpdateCar(car models.Car) (models.Car, error)
	DeleteCar(car models.Car, ID int) (models.Car, error)
}

func RepositoryCar(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindCars() ([]models.Car, error) {
	var cars []models.Car
	err := repo.db.Find(&cars).Error

	return cars, err
}

func (repo *repository) GetCar(ID int) (models.Car, error) {
	var car models.Car
	err := repo.db.First(&car, ID).Error

	return car, err
}

func (repo *repository) CreateCar(car models.Car) (models.Car, error) {
	err := repo.db.Create(&car).Error

	return car, err
}

func (repo *repository) UpdateCar(car models.Car) (models.Car, error) {
	err := repo.db.Save(&car).Error

	return car, err
}

func (repo *repository) DeleteCar(car models.Car, ID int) (models.Car, error) {
	err := repo.db.Delete(&car, ID).Scan(&car).Error

	return car, err
}
