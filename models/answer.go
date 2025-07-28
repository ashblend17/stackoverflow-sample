package models

import "time"

type Answer struct {
	ID         int       `gorm:"column:id;primaryKey"`
	QuestionID int       `gorm:"column:question_id;not null"`
	UserID     int       `gorm:"column:user_id;not null"`
	Question   Question  `gorm:"foreignKey:QuestionID"`
	User       User      `gorm:"foreignKey:UserID"`
	Body       string    `gorm:"column:body;not null"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
