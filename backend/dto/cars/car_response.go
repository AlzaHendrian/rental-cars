package carsdto

import "backend/models"

type ResponseCars struct {
	CarID     int         `json:"car_id"`
	CarName   string      `json:"car_name"`
	DayRate   float64     `json:"day_rate"`
	MonthRate float64     `json:"month_rate"`
	Image     string      `json:"image"`
	UserID    int         `json:"user_id"`
	User      models.User `json:"user"`
}
