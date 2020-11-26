package controllers

// GetPersonByIDRequest
type GetPersonByIDRequest struct {
	IdPerson string `json:"id_person"`
	Age      int    `json:"age"`
}

// GetPersonAddressByIDRequest
type GetPersonAddressByIDRequest struct {
	IdPerson  string `json:"id_person"`
	IdAddress string `json:"id_address"`
}
