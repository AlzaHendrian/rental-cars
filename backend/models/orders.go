package models

import "time"

type Order struct {
	OrderID         int       `gorm:"primaryKey;autoIncrement" json:"order_id"`
	PickupDate      time.Time `gorm:"not null" json:"pickup_date" validate:"required"`
	TimeDropoffDate time.Time `gorm:"not null" json:"dropoff_date" validate:"required"`
	PickupLocation  string    `gorm:"type:varchar(50);not null" json:"pickup_location" validate:"required"`
	DropoffLocation string    `gorm:"type:varchar(50);not null" json:"dropoff_location" validate:"required"`
	Car             Cars      `json:"car"`
	CarID           int       `gorm:"foreignKey;not null" json:"car_id" validate:"required"`
	UserID          int       `gorm:"foreignKey;not null" json:"user_id"`
	User            User      `json:"user"`
}
