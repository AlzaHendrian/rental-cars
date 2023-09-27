package ordersdto

import "time"

type ResponseOrder struct {
	PickupDate      time.Time `json:"pickup_date"`
	TimeDropoffDate time.Time `json:"dropoff_date"`
	PickupLocation  string    `json:"pickup_location"`
	DropoffLocation string    `json:"dropoff_location"`
}
