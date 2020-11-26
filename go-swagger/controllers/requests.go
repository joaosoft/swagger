package controllers

// swagger:parameters GetPersonByIDRequestParameter
type GetPersonByIDRequest struct {
	// in:path
	// id person
	IdPerson string `json:"id_person"`

	// in:query
	// age
	Age int `json:"age"`
}

// swagger:parameters GetPersonAddressByIDRequestParameter
type GetPersonAddressByIDRequest struct {
	// in:path
	// id person
	IdPerson string `json:"id_person"`

	// in:path
	// id address
	IdAddress string `json:"id_address"`
}
