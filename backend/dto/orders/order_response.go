package ordersdto

import (
	"backend/models"
	"time"
)

type ResponseOrder struct {
	PickupDate      time.Time   `json:"pickup_date"`
	TimeDropoffDate time.Time   `json:"dropoff_date"`
	PickupLocation  string      `json:"pickup_location"`
	DropoffLocation string      `json:"dropoff_location"`
	CarID           int         `json:"car_id"`
	Car             models.Cars `json:"car"`
	UserID          int         `json:"user_id"`
	User            models.User `json:"user"`
}
