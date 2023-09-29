package models

type Cars struct {
	CarID     int     `json:"car_id" gorm:"primary_key:auto_increment"`
	CarName   string  `json:"car_name" gorm:"type: varchar(255)"`
	DayRate   float64 `json:"day_rate"`
	MonthRate float64 `json:"month_rate"`
	Image     string  `json:"image" gorm:"type: varchar(255)"`
	UserID    int     `json:"user_id" form:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User      User    `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
