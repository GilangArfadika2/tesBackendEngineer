package controllers

import (
	"GOproject/internal/models"
	"GOproject/internal/services"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Class Controller Customer
type CustomerController struct {
	CustomerService *services.CustomerService
	validator       *validator.Validate
}

func NewCustomerController(service *services.CustomerService, validator *validator.Validate) *CustomerController {
	return &CustomerController{CustomerService: service, validator: validator}
}

// fungsi untuk GET All customer
func (c *CustomerController) GetAllCustomers(ctx *fiber.Ctx) error {
	customers, err := c.CustomerService.GetAllCustomer()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(customers)
}

// fungsi Get customer By ID

func (c *CustomerController) GetCustomerById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	customerId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customer ID tidak valid"})
	}

	customer, err := c.CustomerService.GetCustomerStoreByID(uint(customerId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(customer)
}

// Fungsi menghapus Customer
func (c *CustomerController) DeleteCustomer(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	customerId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customer ID tidak valid"})
	}

	err = c.CustomerService.DeleteCustomer(uint(customerId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Berhasil": "Customer berhasil dihapus"})
}

// fungsi membuat customer
func (c *CustomerController) CreateCustomer(ctx *fiber.Ctx) error {

	var customerStore models.Customer

	if err := ctx.BodyParser(&customerStore); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})

	}

	// Validate the customerStore struct
	if err := validator.New().Struct(customerStore); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	customer, err := c.CustomerService.CreateCustomer(&customerStore)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

// fungsi memperbarui customer
func (c *CustomerController) UpdateCustomer(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	customerId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customer ID tidak valid"})
	}

	var customerStore models.Customer

	if err := ctx.BodyParser(&customerStore); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})

	}

	// Validate the customerStore struct
	if err := validator.New().Struct(customerStore); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	customerStore.ID = uint(customerId)

	customer, err := c.CustomerService.UpdateCustomer(&customerStore)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(customer)
}
