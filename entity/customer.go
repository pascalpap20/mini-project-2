package entity

import "time"

type Customer struct {
	ID        uint   `gorm:"primary_key"`
	FirstName string `gorm:"column:first_name"`
	LastName  string `gorm:"column:last_name"`
	Email     string `gorm:"column:email"`
	Avatar    string `gorm:"column:avatar"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
