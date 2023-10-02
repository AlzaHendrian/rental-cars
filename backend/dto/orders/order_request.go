package ordersdto

import "time"

type CreateRequestOrder struct {
	UserID          int       `json:"user_id" form:"user_id" validate:"required"`
	CarID           int       `json:"car_id" form:"car_id" validate:"required"`
	PickupDate      time.Time `json:"pickup_date" form:"pickup_date" validate:"required"`
	TimeDropoffDate time.Time `json:"dropoff_date" form:"dropoff_date" validate:"required"`
	PickupLocation  string    `json:"pickup_location" form:"pickup_location" validate:"required"`
	DropoffLocation string    `json:"dropoff_location" form:"dropoff_location" validate:"required"`
}
