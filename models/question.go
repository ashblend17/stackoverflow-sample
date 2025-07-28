package models

import "time"

type Question struct {
	ID        int       `gorm:"column:id;primaryKey"`
	UserID    int       `gorm:"column:user_id;not null"`
	User      User      `gorm:"foreignKey:UserID"`
	Title     string    `gorm:"column:title;not null"`
	Body      string    `gorm:"column:body;not null"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
