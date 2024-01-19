package response

import "GOproject/internal/models"

// dto untuk memberikan response data customer
type CustomerResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
	Email       string `json:"email"`
}

func GetAllCustomerResponse(listCustomer *[]models.Customer) *[]CustomerResponse {

	var listResponse []CustomerResponse

	for _, customer := range *listCustomer {
		respon := &CustomerResponse{
			ID:          customer.ID,
			Title:       customer.Title,
			Name:        customer.Name,
			Gender:      customer.Gender,
			PhoneNumber: customer.PhoneNumber,
			Image:       customer.Image,
			Email:       customer.Email,
		}
		listResponse = append(listResponse, *respon)
	}

	return &listResponse
}

func GetCustomerResponse(customer *models.Customer) *CustomerResponse {

	respon := &CustomerResponse{
		ID:          customer.ID,
		Title:       customer.Title,
		Name:        customer.Name,
		Gender:      customer.Gender,
		PhoneNumber: customer.PhoneNumber,
		Image:       customer.Image,
		Email:       customer.Email,
	}

	return respon
}
