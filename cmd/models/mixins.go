package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID 		  uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	now := time.Now().UTC()
	b.ID = id
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

func (b *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	b.UpdatedAt = time.Now().UTC()
	return nil
}