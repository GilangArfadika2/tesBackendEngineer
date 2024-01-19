package response

import "GOproject/internal/models"

// dto untuk memberikan detai data customer
type CustomerResponseDetail struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Name        string            `json:"name"`
	Gender      string            `json:"gender"`
	PhoneNumber string            `json:"phone_number"`
	Image       string            `json:"image"`
	Email       string            `json:"email"`
	Address     []AddressResponse `json:"address"`
}

func GetCustomerRenponseById(customer *models.Customer) *CustomerResponseDetail {

	var addressResponse []AddressResponse

	for _, address := range customer.Address {
		respon := CreateAddressResponse(&address)
		addressResponse = append(addressResponse, *respon)
	}

	response := &CustomerResponseDetail{
		ID:          customer.ID,
		Title:       customer.Title,
		Name:        customer.Name,
		Gender:      customer.Gender,
		PhoneNumber: customer.PhoneNumber,
		Image:       customer.Image,
		Email:       customer.Email,
		Address:     addressResponse,
	}

	return response

}
