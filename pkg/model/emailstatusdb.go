package model

import "time"

type EmailStatusGORM struct {
	ID        int64 `gorm:"primaryKey"`
	Status    string
	Typ       string `gorm:"column:type"`
	SentAt    time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// override table name with email_status
func (EmailStatusGORM) TableName() string {
	return "email_status"
}
