package dto

import "github.com/markaz/models"

type OrderRequest struct {
	UserId            int64                  `json:"user_id"`
	ShippingAddress   Address                `json:"shipping_address"`
	BillingAddress    Address                `json:"billing_address"`
	Products          []models.Product       `json:"products"`
	PaymentMethod     string                 `json:"payment_method"`
	PaymentMethodData map[string]interface{} `json:"payment_method_data"`
}

type Address struct {
	Address1    string `json:"address_1"`
	Address2    string `json:"address_2,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number,omitempty"`
	State       string `json:"state,omitempty"`
	Zip         string `json:"zip"`
}
