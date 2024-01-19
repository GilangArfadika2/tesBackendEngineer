package seeders

import (
	"GOproject/internal/models"
	"GOproject/internal/repository"
	"strconv"

	"gorm.io/gorm"
)

func SeedCustomer(db *gorm.DB) error {

	customerDb := repository.NewGormCustomerStoreRepository(db)

	// listCustomer := []models.Customer{
	// 	{Title: "Mr", Name: "John Doe", Gender: "M", PhoneNumber: "123456789", Image: "john.jpg", Email: "john@example.com"},
	// 	{Title: "Ms", Name: "Jane Doe", Gender: "F", PhoneNumber: "987654321", Image: "jane.jpg", Email: "jane@example.com"},
	// }
	var listCustomer []models.Customer

	for i := 0; i < 10; i++ {
		cust := models.Customer{
			Title:       "Title" + strconv.Itoa(i),
			Name:        "User " + strconv.Itoa(i),
			Gender:      "M",
			PhoneNumber: "123456" + strconv.Itoa(i),
			Image:       "user" + strconv.Itoa(i) + ".jpg",
			Email:       "user" + strconv.Itoa(i) + "@example.com",
		}

		listCustomer = append(listCustomer, cust)
	}

	for _, customer := range listCustomer {
		_, err := customerDb.CreateCustomerStore(&customer)
		if err != nil {
			return err
		}
	}

	return nil

}
