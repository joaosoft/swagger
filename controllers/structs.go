package controllers

// A SuccessResponse is an error that is used when the required input fails validation.
// swagger:response SuccessResponse
type SuccessResponse struct {
	// The validation message
	//
	// Required: true
	Success bool `json:"success"`
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
