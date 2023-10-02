package models

type User struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment; type:integer;not null"`
	Name     string `json:"name" gorm:"type: varchar(255);not null"`
	Email    string `json:"email" gorm:"type: varchar(255);not null"`
	Password string `json:"password" gorm:"type: varchar(255);not null"`
}
