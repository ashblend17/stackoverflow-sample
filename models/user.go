package models

import "time"

type User struct {
	ID        int       `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username;unique;not null"`
	Email     string    `gorm:"column:email;unique;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
