package model

type ContactResponse struct {
	ID        string            `json:"id" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	FirstName string            `json:"first_name" example:"John"`
	LastName  string            `json:"last_name" example:"Doe"`
	Email     string            `json:"email" example:"john.doe@example.com"`
	Phone     string            `json:"phone" example:"081234567890"`
	CreatedAt int64             `json:"created_at" example:"1630000000"`
	UpdatedAt int64             `json:"updated_at" example:"1630000000"`
	Addresses []AddressResponse `json:"addresses,omitempty" example:"[]"`
}

type CreateContactRequest struct {
	UserId    string `json:"-" validate:"required" example:"user1"`
	FirstName string `json:"first_name" validate:"required,max=100" example:"John"`
	LastName  string `json:"last_name" validate:"max=100" example:"Doe"`
	Email     string `json:"email" validate:"max=200,email" example:"john.doe@example.com"`
	Phone     string `json:"phone" validate:"max=20" example:"081234567890"`
}

type UpdateContactRequest struct {
	UserId    string `json:"-" validate:"required" example:"user1"`
	ID        string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
	FirstName string `json:"first_name" validate:"required,max=100" example:"John"`
	LastName  string `json:"last_name" validate:"max=100" example:"Doe"`
	Email     string `json:"email" validate:"max=200,email" example:"john.doe@example.com"`
	Phone     string `json:"phone" validate:"max=20" example:"081234567890"`
}

type SearchContactRequest struct {
	UserId string `json:"-" validate:"required" example:"user1"`
	Name   string `json:"name" validate:"max=100" example:"John Doe"`
	Email  string `json:"email" validate:"max=200" example:"john.doe@example.com"`
	Phone  string `json:"phone" validate:"max=20" example:"081234567890"`
	Page   int    `json:"page" validate:"min=1" example:"1"`
	Size   int    `json:"size" validate:"min=1,max=100" example:"10"`
}

type GetContactRequest struct {
	UserId string `json:"-" validate:"required" example:"user1"`
	ID     string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
}

type DeleteContactRequest struct {
	UserId string `json:"-" validate:"required" example:"user1"`
	ID     string `json:"-" validate:"required,max=100,uuid" example:"013b4e-78ac-73dc-8480-da4ec118ef28"`
}
