// repository/customer_store_repository.go
package repository

import (
	"GOproject/internal/models"

	"gorm.io/gorm"
)

// interface untuk repository data Addresss
type AddressRepository interface {
	GetAllAddress() (*[]models.Address, error)
	GetAddressById(uint) (*models.Address, error)
	CreateAddress(*models.Address, uint, CustomerStoreRepository) (*models.Address, error)

	UpdateAddress(*models.Address) (*models.Address, error)
	DeleteAddress(uint) error
}

// implementasi GORM untuk repository
type GormAddressRepository struct {
	db *gorm.DB
}

// NewGormAddressRepository creates a new instance of GormAddressRepository.
func NewGormAddressRepository(db *gorm.DB) AddressRepository {
	return &GormAddressRepository{db: db}
}

// mendapatkan semua address di Db
func (r *GormAddressRepository) GetAllAddress() (address *[]models.Address, err error) {

	var addressList []models.Address
	if err := r.db.Find(&addressList).Error; err != nil {
		return nil, err
	}
	return &addressList, nil
}

// membuat address berdasarkan id customer
func (r *GormAddressRepository) CreateAddress(address *models.Address, id uint, r2 CustomerStoreRepository) (addressNew *models.Address, err error) {

	customer, err := r2.GetCustomerStoreByID(id)
	if err != nil {
		return nil, err
	}

	address.CustomerID = *&customer.ID

	result := r.db.Create(address)

	if result.Error != nil {
		return nil, result.Error
	}

	customer.Address = append(customer.Address, *address)

	return address, nil

}

// mendapatkan customer by id
func (r *GormAddressRepository) GetAddressById(id uint) (*models.Address, error) {
	var address models.Address
	err := r.db.First(&address, id).Error
	return &address, err
}

// memperbarui address di Db
func (r *GormAddressRepository) UpdateAddress(address *models.Address) (*models.Address, error) {

	addressOld, err := r.GetAddressById(address.ID)

	if err != nil {
		return nil, err
	}

	address.CreatedAt = addressOld.CreatedAt
	address.CustomerID = addressOld.CustomerID

	result := r.db.Save(address)

	if result.Error != nil {
		return nil, result.Error
	}

	return address, nil
}

// DeleteCustomerStore deletes a customer store by ID from the database.
func (r *GormAddressRepository) DeleteAddress(id uint) error {

	address, err := r.GetAddressById(id)

	if err != nil {
		return err
	}
	return r.db.Delete(address).Error
}
