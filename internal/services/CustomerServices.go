package services

import (
	"GOproject/internal/dto/response"
	"GOproject/internal/models"
	"GOproject/internal/repository"
)

type CustomerService struct {
	customerDb repository.CustomerStoreRepository
}

func NewCustomerService(customerDb repository.CustomerStoreRepository) *CustomerService {

	return &CustomerService{customerDb: customerDb}
}

func (customerService *CustomerService) GetAllCustomer() (customerStore *[]response.CustomerResponse, err error) {

	customerlist, err := customerService.customerDb.GetAllCustomerStore()

	response := response.GetAllCustomerResponse(customerlist)

	return response, err

}

func (customerService *CustomerService) CreateCustomer(customer *models.Customer) (*response.CustomerResponse, error) {

	customerSaved, err := customerService.customerDb.CreateCustomerStore(customer)

	respon := response.GetCustomerResponse(customerSaved)

	if err != nil {
		return nil, err
	}

	return respon, nil

}

func (customerService *CustomerService) UpdateCustomer(customer *models.Customer) (*response.CustomerResponse, error) {

	customerSaved, err := customerService.customerDb.UpdateCustomerStore(customer)

	if err != nil {
		return nil, err
	}

	respon := response.GetCustomerResponse(customerSaved)

	return respon, nil

}

func (customerService *CustomerService) GetCustomerStoreByID(id uint) (*response.CustomerResponseDetail, error) {

	customer, err := customerService.customerDb.GetCustomerStoreByID(id)

	if err != nil {
		return nil, err
	}

	respon := response.GetCustomerRenponseById(customer)

	return respon, nil
}

func (customerService *CustomerService) DeleteCustomer(id uint) error {

	return customerService.customerDb.DeleteCustomerStore(id)
}

// func CreateCustomer(customerRepo repository.CustomerStoreRepository, customer *models.Customer) error {

// 	// Create the customer in the database
// 	if err := customerRepo.CreateCustomerStore(customer); err != nil {
// 		return err
// 	}

// 	return nil

// }
