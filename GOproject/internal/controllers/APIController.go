package controllers

// import (
// 	"github.com/go-playground/validator/v10"
// )

// // Class Controller Customer
// type APIController struct {
// 	validator *validator.Validate
// }

// func NewAPIController(validator *validator.Validate) *APIController {
// 	return &APIController{validator: validator}
// }


// // fungsi untuk GET All customer
// func (c *CustomerController) GetAllCustomers(ctx *fiber.Ctx) error {

// 	if err != nil {
// 		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
// 	}
// 	return ctx.Status(fiber.StatusOK).JSON(customers)
// }