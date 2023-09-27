package carsdto

type CreateCar struct {
	CarName   string  `json:"car_name" form:"car_name" validate:"required"`
	DayRate   float64 `json:"day_rate" form:"day_rate" validate:"required"`
	MonthRate float64 `json:"month_rate" form:"month_rate" validate:"required"`
	Image     string  `json:"image" form:"image" validate:"required"`
}
