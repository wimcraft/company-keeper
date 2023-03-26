package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key"`
	Name         string    `gorm:"type:varchar(15);uniqueIndex;not null"`
	Description  string    `gorm:"type:text"`
	NumEmployees int       `gorm:"not null"`
	Registered   bool      `gorm:"not null"`
	Type         string    `gorm:"type:varchar(20);not null"`
}

func (c *Company) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New()
	return nil
}
