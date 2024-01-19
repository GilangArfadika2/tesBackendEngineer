package models

import (
	"time"

	"gorm.io/gorm"
)

// CustomerStore represents the model for a customer store.
// CustomerStore represents the model for a customer store.
type Customer struct {
	ID          uint       `json:"id" `
	Title       string     `json:"title" validate:"required"`
	Name        string     `json:"name" validate:"required"`
	Gender      string     `json:"gender" validate:"required"`
	PhoneNumber string     `json:"phone_number" validate:"required"`
	Image       string     `json:"image" validate:"required"`
	Email       string     `json:"email" validate:"required"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	Address     []Address  `json:"address" gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	customer.CreatedAt = &now
	customer.UpdatedAt = nil
	return nil
}

func (customer *Customer) BeforeUpdate(tx *gorm.DB) (err error) {

	now := time.Now()
	customer.UpdatedAt = &now

	return nil
}
