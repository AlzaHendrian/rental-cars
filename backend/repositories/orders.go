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
	err := repo.db.Preload("Car").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email").Table("users")
	}).Find(&orders).Error

	return orders, err
}

func (r *repository) FindOrderByUser(UserID int) ([]models.Order, error) {
	var transaction []models.Order
	err := r.db.Preload("Car").Where("user_id = ?", UserID).Error

	return transaction, err
}

func (repo *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := repo.db.Preload("Car").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "name", "email").Table("users")
	}).First(&order, ID).Error

	return order, err
}

func (repo *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := repo.db.Preload("Car").Create(&order).Error

	return order, err
}

func (repo *repository) UpdateOrder(order models.Order) (models.Order, error) {
	err := repo.db.Preload("Car").Save(&order).Error

	return order, err
}

func (repo *repository) DeleteOrder(order models.Order, ID int) (models.Order, error) {
	err := repo.db.Delete(&order, ID).Scan(&order).Error

	return order, err
}
