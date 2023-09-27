package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type CarRepository interface {
	FindCars() ([]models.Cars, error)
	GetCar(ID int) (models.Cars, error)
	CreateCar(car models.Cars) (models.Cars, error)
	UpdateCar(car models.Cars) (models.Cars, error)
	DeleteCar(car models.Cars, ID int) (models.Cars, error)
}

func RepositoryCar(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindCars() ([]models.Cars, error) {
	var cars []models.Cars
	err := repo.db.Find(&cars).Error

	return cars, err
}

func (repo *repository) GetCar(ID int) (models.Cars, error) {
	var car models.Cars
	err := repo.db.First(&car, ID).Error

	return car, err
}

func (repo *repository) CreateCar(car models.Cars) (models.Cars, error) {
	err := repo.db.Create(&car).Error

	return car, err
}

func (repo *repository) UpdateCar(car models.Cars) (models.Cars, error) {
	err := repo.db.Save(&car).Error

	return car, err
}

func (repo *repository) DeleteCar(car models.Cars, ID int) (models.Cars, error) {
	err := repo.db.Delete(&car, ID).Scan(&car).Error

	return car, err
}
