package models

import "time"

type User struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UserId    int       `gorm:"column:user_id;PRIMARY_KEY"`
	Name      string    `gorm:"column:name"`
}
