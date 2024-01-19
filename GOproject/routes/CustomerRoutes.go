package routes

import (
	"GOproject/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoutes(app *fiber.App, customerController *controllers.CustomerController) {

	// Define your routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Define your routes
	app.Get("/customer-store", customerController.GetAllCustomers)

	// Route for creating a new customer (POST request)
	app.Post("/customer-store", customerController.CreateCustomer)

	// Route for getting a customer by Id using Path Variable(GET request)
	app.Get("/customer-store/:id", customerController.GetCustomerById)

	// Route for updating a customer by Id using Path Variable and Request Body (PATCH request)
	app.Patch("/customer-store/:id", customerController.UpdateCustomer)

	// Route for Deleting a customer by Id using Path Variable(Delete request)
	app.Delete("/customer-store/:id", customerController.DeleteCustomer)

}
