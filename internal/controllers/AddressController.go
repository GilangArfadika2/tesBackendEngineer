package controllers

import (
	"GOproject/internal/models"
	"GOproject/internal/services"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AddressController struct {
	addressService *services.AddressService
	validator      *validator.Validate
}

func NewAddressController(service *services.AddressService, validate *validator.Validate) *AddressController {
	return &AddressController{addressService: service, validator: validate}
}

func (c *AddressController) GetAllAddress(ctx *fiber.Ctx) error {
	address, err := c.addressService.GetAllAddress()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(address)
}

func (c *AddressController) GetAddressById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	addressId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "address ID tidak valid"})
	}

	address, err := c.addressService.GetAddressById(uint(addressId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(address)
}

// DELETE HARUS DIPERBAIKI
func (c *AddressController) DeleteAddress(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	addressId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "address ID tidak valid"})
	}

	err = c.addressService.DeleteAddress(uint(addressId))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"Berhasil": "address berhasil dihapus"})
}

func (c *AddressController) CreateAddress(ctx *fiber.Ctx) error {

	var address models.Address

	id := ctx.Params("id")

	customerId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "customer ID tidak valid"})
	}

	if err := ctx.BodyParser(&address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})

	}

	// Validate the customerStore struct
	if err := validator.New().Struct(address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	addressSaved, err := c.addressService.CreateAddress(&address, uint(customerId))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return ctx.Status(fiber.StatusOK).JSON(addressSaved)
}

func (c *AddressController) UpdateAddress(ctx *fiber.Ctx) error {

	id := ctx.Params("id")

	addressId, err := strconv.Atoi(id)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Address ID tidak valid"})
	}

	var address models.Address

	if err := ctx.BodyParser(&address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Bad Request"})

	}
	// Validate the customerStore struct
	if err := validator.New().Struct(address); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	address.ID = uint(addressId)

	addressUpdated, err := c.addressService.UpdateAddress(&address)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return ctx.Status(fiber.StatusOK).JSON(addressUpdated)
}
