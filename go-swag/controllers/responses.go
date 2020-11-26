package controllers

// PersonResponse is an struct that contains the person information.
type PersonResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// AddressResponse is an struct that contains the address information.
type AddressResponse struct {
	Id      string `json:"id"`
	Country string `json:"country"`
	City    string `json:"city"`
	Street  string `json:"street"`
	Number  int    `json:"number"`
}

// ErrorResponse is an error that is used when the required input fails validation.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
