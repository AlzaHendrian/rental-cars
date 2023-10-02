package models

type Car struct {
	CarID     int     `json:"car_id" gorm:"primary_key:auto_increment; type:integer;not null"`
	CarName   string  `json:"car_name" gorm:"type: varchar(50);not null"`
	DayRate   float64 `json:"day_rate" gorm:"not null"`
	MonthRate float64 `json:"month_rate" gorm:"not null"`
	Image     string  `json:"image" gorm:"type: varchar(256);not null"`
}
