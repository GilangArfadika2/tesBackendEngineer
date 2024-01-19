package seeders

import (
	"GOproject/internal/models"
	"GOproject/internal/repository"
	"strconv"

	"gorm.io/gorm"
)

func SeedAddress(db *gorm.DB) error {

	addressDb := repository.NewGormAddressRepository(db)
	customerDb := repository.NewGormCustomerStoreRepository(db)

	// listAddress := []models.Address{
	// 	{
	//
	// 		Address:    "123 Main St",
	// 		District:   "Downtown",
	// 		City:       "Cityville",
	// 		Province:   "State",
	// 		PostalCode: 12345,
	// 	},
	// 	{
	//
	// 		Address:    "456 Oak St",
	// 		District:   "Uptown",
	// 		City:       "Townsville",
	// 		Province:   "State",
	// 		PostalCode: 56789,
	// 	},
	// 	// Add more addresses as needed
	// }

	var listAddress []models.Address
	for i := 1; i < 10; i++ {
		addrs := models.Address{
			CustomerID: uint(i),
			Address:    "123 Main St " + strconv.Itoa(i),
			District:   "Downtown " + strconv.Itoa(i),
			City:       "Cityville " + strconv.Itoa(i),
			Province:   "State " + strconv.Itoa(i),
			PostalCode: 12345 + i,
		}

		listAddress = append(listAddress, addrs)
	}

	for _, address := range listAddress {
		if _, err := addressDb.CreateAddress(&address, address.CustomerID, customerDb); err != nil {
			return err
		}
	}

	// for _, customer := range listCustomer {
	// 	_, err := customerDb.CreateCustomerStore(&customer)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil

}
