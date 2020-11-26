package controllers

// A PersonResponse is an struct that contains the person information.
// swagger:response PersonResponse
type PersonResponse struct {
	// Id
	Id string `json:"id"`

	// Name
	Name string `json:"name"`

	// Age
	Age int `json:"age"`
}

// A AddressResponse is an struct that contains the address information.
// swagger:response AddressResponse
type AddressResponse struct {
	// Id
	Id string `json:"id"`

	// Country
	Country string `json:"country"`

	// City
	City string `json:"city"`

	// Street
	Street string `json:"street"`

	// Number
	Number int `json:"number"`
}

// A ErrorResponse is an error that is used when the required input fails validation.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// Code
	Code int `json:"code"`
	// Message
	Message string `json:"message"`
}
