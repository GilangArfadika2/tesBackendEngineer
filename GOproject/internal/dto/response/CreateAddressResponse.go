package response

import (
	"GOproject/internal/models"
	"time"
)

// Dto untuk memberikan response address
type AddressResponse struct {
	ID         uint       `json:"id"`
	CustomerID uint       `json:"customer_id"`
	Address    string     `json:"address"`
	District   string     `json:"district"`
	City       string     `json:"city"`
	Province   string     `json:"province"`
	PostalCode int        `json:"postal_code"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

func CreateAddressResponse(address *models.Address) *AddressResponse {

	response := &AddressResponse{
		ID:         address.ID,
		CustomerID: address.CustomerID,
		Address:    address.Address,
		District:   address.District,
		City:       address.City,
		Province:   address.Province,
		PostalCode: address.PostalCode,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}

	return response

}
