package controllers

// A PersonResponse is an struct that contains the person information.
// swagger:response PersonResponse
type PersonResponse struct {
	// The validation message
	//
	// Required: true
	Id string `json:"id"`

	// Required: true
	Name string `json:"name"`

	// Required: true
	Age int `json:"age"`
}

// A AddressResponse is an struct that contains the address information.
// swagger:response AddressResponse
type AddressResponse struct {
	// The validation message
	//
	// Required: true
	Country string `json:"country"`

	// Required: true
	City string `json:"city"`

	// Required: true
	Street string `json:"street"`

	// Required: true
	Number int `json:"number"`
}

// A ErrorResponse is an error that is used when the required input fails validation.
// swagger:response ErrorResponse
type ErrorResponse struct {
	// An optional field code to which this validation applies
	Code int `json:"code"`
	// The validation message
	//
	// Required: true
	// Example: Expected type string
	Message string `json:"message"`
}
