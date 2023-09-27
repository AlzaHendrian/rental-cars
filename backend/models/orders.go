package models

import "time"

type Order struct {
	OrderID         uint      `gorm:"primaryKey;autoIncrement" json:"order_id"`
	CarID           uint      `gorm:"foreignKey;not null" json:"car_id" validate:"required"`
	PickupDate      time.Time `gorm:"not null" json:"pickup_date" validate:"required"`
	TimeDropoffDate time.Time `gorm:"not null" json:"dropoff_date" validate:"required"`
	PickupLocation  string    `gorm:"type:varchar(50);not null" json:"pickup_location" validate:"required"`
	DropoffLocation string    `gorm:"type:varchar(50);not null" json:"dropoff_location" validate:"required"`
}
