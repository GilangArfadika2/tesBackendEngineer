package routes

import (
	"GOproject/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(app *fiber.App, addressController *controllers.AddressController) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Route for Getting all customeer (Get request)
	app.Get("/address", addressController.GetAllAddress)

	// Route for creating a new customer (POST request)
	app.Post("/address/:id", addressController.CreateAddress)

	// Route for getting a customer by Id using Path Variable(GET request)
	app.Get("/address/:id", addressController.GetAddressById)

	// Route for updating a customer by Id using Path Variable and Request Body (PATCH request)
	app.Patch("/address/:id", addressController.UpdateAddress)

	// Route for Deleting a customer by Id using Path Variable(Delete request)
	app.Delete("/address/:id", addressController.DeleteAddress)

}
