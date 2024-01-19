package services

import (
	"GOproject/internal/dto/response"
	"GOproject/internal/models"
	"GOproject/internal/repository"
)

type AddressService struct {
	addressDb repository.AddressRepository

	customerDb repository.CustomerStoreRepository
}

func NewAddressService(db repository.AddressRepository, custDb repository.CustomerStoreRepository) AddressService {

	return AddressService{addressDb: db, customerDb: custDb}

}

func (addressService *AddressService) GetAllAddress() (*[]response.AddressResponse, error) {

	var listResponse []response.AddressResponse
	listAddress, err := addressService.addressDb.GetAllAddress()

	for _, address := range *listAddress {
		respon := response.CreateAddressResponse(&address)
		listResponse = append(listResponse, *respon)
	}

	return &listResponse, err
}

func (addressService *AddressService) CreateAddress(address *models.Address, id uint) (*response.AddressResponse, error) {

	addressSaved, err := addressService.addressDb.CreateAddress(address, id, addressService.customerDb)

	if err != nil {
		return nil, err
	}

	// if addressSaved.Customer != nil {
	// 	addressSaved.Customer = nil
	// }
	// Create an AddressResponse object
	response := response.CreateAddressResponse(addressSaved)

	return response, nil
}

func (addressService *AddressService) GetAddressById(id uint) (*response.AddressResponse, error) {

	address, err := addressService.addressDb.GetAddressById(id)

	if err != nil {
		return nil, err
	}

	response := response.CreateAddressResponse(address)

	return response, nil
}

func (addressService *AddressService) UpdateAddress(address *models.Address) (*response.AddressResponse, error) {

	addressSaved, err := addressService.addressDb.UpdateAddress(address)

	if err != nil {
		return nil, err
	}

	response := response.CreateAddressResponse(addressSaved)

	return response, nil

}

func (addressService *AddressService) DeleteAddress(id uint) error {
	err := addressService.addressDb.DeleteAddress(id)

	if err != nil {
		return err
	}

	return nil
}
