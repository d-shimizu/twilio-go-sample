package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type PhoneNumber struct {
	ID           uint            `gorm:"primaryKey"`
	TwilioSid    string          `gorm:"type:varchar(34);unique;not null"`
	PhoneNumber  string          `gorm:"type:varchar(20);unique;not null"`
	AccountSid   string          `gorm:"type:varchar(34);not null"`
	Status       string          `gorm:"type:varchar(20);not null;default:'active'"`
	Capabilities json.RawMessage `gorm:"type:jsonb"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
