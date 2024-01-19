// repository/customer_store_repository.go
package repository

import (
	"GOproject/internal/models"

	"gorm.io/gorm"
)

// Interface untuk mengatur repository customer
type CustomerStoreRepository interface {
	CreateCustomerStore(customerStore *models.Customer) (*models.Customer, error)
	GetCustomerStoreByID(id uint) (*models.Customer, error)
	UpdateCustomerStore(customerStore *models.Customer) (*models.Customer, error)
	DeleteCustomerStore(id uint) error
	GetAllCustomerStore() (*[]models.Customer, error)
	GetDB() *gorm.DB
}

// implementasi GORM untuk repository customer
type GormCustomerStoreRepository struct {
	db *gorm.DB
}

// NewGormCustomerStoreRepository creates a new instance of GormCustomerStoreRepository.
func NewGormCustomerStoreRepository(db *gorm.DB) CustomerStoreRepository {
	return &GormCustomerStoreRepository{db: db}
}

// Mendapatkan semua customer
func (r *GormCustomerStoreRepository) GetAllCustomerStore() (customerStore *[]models.Customer, err error) {

	var customerLis []models.Customer
	if err := r.db.Find(&customerLis).Error; err != nil {
		return nil, err
	}
	return &customerLis, nil
}

// membuat customer
func (r *GormCustomerStoreRepository) CreateCustomerStore(customerStore *models.Customer) (customer *models.Customer, err error) {

	result := r.db.Create(customerStore)

	if result.Error != nil {
		return nil, result.Error
	}

	return customerStore, nil
}

func (r *GormCustomerStoreRepository) GetDB() *gorm.DB {
	return r.db
}

// mendapatkan customer by id .
func (r *GormCustomerStoreRepository) GetCustomerStoreByID(id uint) (*models.Customer, error) {
	var customerStore models.Customer
	// Preload the associated addresses
	result := r.db.Preload("Address").First(&customerStore, id)

	if result.Error != nil {
		return nil, result.Error
	}

	err := r.db.First(&customerStore, id).Error
	return &customerStore, err
}

// memperbarui customer by id customer
func (r *GormCustomerStoreRepository) UpdateCustomerStore(customerStore *models.Customer) (*models.Customer, error) {

	customer, err := r.GetCustomerStoreByID(customerStore.ID)

	if err != nil {
		return nil, err
	}

	customerStore.CreatedAt = customer.CreatedAt

	result := r.db.Save(customerStore)

	if result.Error != nil {
		return nil, result.Error
	}

	return customerStore, nil
}

// menghapus customer dari repository
func (r *GormCustomerStoreRepository) DeleteCustomerStore(id uint) error {

	customer, err := r.GetCustomerStoreByID(id)

	if err != nil {
		return err
	}
	return r.db.Delete(customer).Error
}
