package models

import (
	"time"

	"gorm.io/gorm"
)

// Address represents the model for an address.
type Address struct {
	ID         uint       `json:"id"`
	CustomerID uint       `json:"customer_id"`
	Customer   *Customer  `json:"customer" gorm:"foreignKey:CustomerID"`
	Address    string     `json:"address" validate:"required"`
	District   string     `json:"district" validate:"required"`
	City       string     `json:"city" validate:"required"`
	Province   string     `json:"province" validate:"required"`
	PostalCode int        `json:"postal_code" validate:"required"`
	CreatedAt  *time.Time `json:"created_at" `
	UpdatedAt  *time.Time `json:"updated_at"`
}

// BeforeCreate is a GORM hook to set the created_at timestamp before creating a new record.
func (address *Address) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	address.CreatedAt = &now
	address.UpdatedAt = nil
	return nil
}

// BeforeUpdate is a GORM hook to set the updated_at timestamp before updating a record.
func (address *Address) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	address.UpdatedAt = &now
	return nil
}
