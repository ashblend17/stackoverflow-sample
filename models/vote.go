package models

import "time"

type Vote struct {
	ID        int       `gorm:"column:id;primaryKey"`
	UserID    int       `gorm:"column:user_id;not null"`
	ItemID    int       `gorm:"column:item_id;not null"`
	ItemType  string    `gorm:"column:item_type;not null"` // 'question' or 'answer'
	VoteType  string    `gorm:"column:vote_type;not null"` // 'upvote' or 'downvote'
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
