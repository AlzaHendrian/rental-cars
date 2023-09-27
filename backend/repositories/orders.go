package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order, ID int) (models.Order, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order
	err := repo.db.Find(&orders).Error

	return orders, err
}

func (repo *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := repo.db.First(&order, ID).Error

	return order, err
}

func (repo *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := repo.db.Create(&order).Error

	return order, err
}

func (repo *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := repo.db.Save(&order).Error

	return order, err
}

func (repo *repository) DeleteOrder(order models.Order, ID int) (models.Order, error) {
	err := repo.db.Delete(&order, ID).Scan(&order).Error

	return order, err
}
