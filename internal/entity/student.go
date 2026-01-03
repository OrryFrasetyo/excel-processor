package entity

import "time"

type Student struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	CreatedAt time.Time
}
