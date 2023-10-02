package models

import "time"

type Order struct {
	OrderID         int       `gorm:"primaryKey;autoIncrement; type:integer;not null" json:"order_id"`
	CarID           int       `gorm:"foreignKey:car_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"car_id"`
	Car             Car       `json:"car"`
	PickupDate      time.Time `gorm:"type:date;not null" json:"pickup_date" validate:"required"`
	TimeDropoffDate time.Time `gorm:"type:date;not null" json:"dropoff_date" validate:"required"`
	PickupLocation  string    `gorm:"type:varchar(50);not null" json:"pickup_location" validate:"required"`
	DropoffLocation string    `gorm:"type:varchar(50);not null" json:"dropoff_location" validate:"required"`
	UserID          int       `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User            User      `json:"user"`
}
