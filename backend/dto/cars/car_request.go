package carsdto

type CreateCar struct {
	CarID     int     `json:"car_id" form:"car_id"`
	CarName   string  `json:"car_name" form:"car_name" validate:"required"`
	DayRate   float64 `json:"day_rate" form:"day_rate" validate:"required"`
	MonthRate float64 `json:"month_rate" form:"month_rate" validate:"required"`
	Image     string  `json:"image" form:"image" validate:"required"`
	UserID    int     `json:"user_id" form:"user_id"`
}
