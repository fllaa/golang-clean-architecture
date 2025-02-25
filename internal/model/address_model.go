package model

type AddressResponse struct {
	ID         string `json:"id"`
	Street     string `json:"street"`
	City       string `json:"city"`
	Province   string `json:"province"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}

type ListAddressRequest struct {
	UserId    string `json:"-" validate:"required" example:"user1"`
	ContactId string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
}

type CreateAddressRequest struct {
	UserId     string `json:"-" validate:"required" example:"user1"`
	ContactId  string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	Street     string `json:"street" validate:"max=255" example:"Jl. Jend. Sudirman No. 1"`
	City       string `json:"city" validate:"max=255" example:"Jakarta"`
	Province   string `json:"province" validate:"max=255" example:"DKI Jakarta"`
	PostalCode string `json:"postal_code" validate:"max=10" example:"10000"`
	Country    string `json:"country" validate:"max=100" example:"Indonesia"`
}

type UpdateAddressRequest struct {
	UserId     string `json:"-" validate:"required" example:"user1"`
	ContactId  string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	ID         string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	Street     string `json:"street" validate:"max=255" example:"Jl. Jend. Sudirman No. 1"`
	City       string `json:"city" validate:"max=255" example:"Jakarta"`
	Province   string `json:"province" validate:"max=255" example:"DKI Jakarta"`
	PostalCode string `json:"postal_code" validate:"max=10" example:"10000"`
	Country    string `json:"country" validate:"max=100" example:"Indonesia"`
}

type GetAddressRequest struct {
	UserId    string `json:"-" validate:"required" example:"user1"`
	ContactId string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	ID        string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
}

type DeleteAddressRequest struct {
	UserId    string `json:"-" validate:"required" example:"user1"`
	ContactId string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	ID        string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
}
