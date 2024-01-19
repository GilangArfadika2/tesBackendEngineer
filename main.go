// main.go
package main

import (
	"GOproject/database/migrations"
	"GOproject/database/seeders"
	"GOproject/routes"
	"log"

	"GOproject/internal/controllers"
	"GOproject/internal/models"
	"GOproject/internal/repository"
	"GOproject/internal/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Menginisiasi fiber
	app := fiber.New()

	// Menginisiasi DB
	db, err := migrations.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}

	// migrasi DB
	err = migrations.MigrateDB(db)
	if err != nil {
		log.Fatal(err)
	}

	// // Reset records in the customers table
	// err = db.Exec("DELETE FROM customers").Error
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Menjalankan seeder ketika table customer pertamaa kali dibuat
	var customerCount int64
	db.Model(&models.Customer{}).Count(&customerCount)

	// Run seeders only if the customers table is empty
	if customerCount == 0 {
		err = seeders.SeedCustomer(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	var AddressCount int64

	// Menjalankan seeder ketika table address pertama kali dibuat
	db.Model(&models.Address{}).Count(&AddressCount)
	if customerCount == 0 {
		err = seeders.SeedAddress(db)
		if err != nil {
			log.Fatal(err)
		}
	}

	//membuat service dan controller
	customerDb := repository.NewGormCustomerStoreRepository(db)
	customerService := services.NewCustomerService(customerDb)
	customerController := controllers.NewCustomerController(customerService, &validator.Validate{})
	addressDb := repository.NewGormAddressRepository(db)
	addressService := services.NewAddressService(addressDb, customerDb)
	addressControler := controllers.NewAddressController(&addressService, &validator.Validate{})

	// Define your routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.CustomerRoutes(app, customerController)

	routes.AddressRoutes(app, addressControler)

	// Start the Fiber app
	err = app.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}

}
