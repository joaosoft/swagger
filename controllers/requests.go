package controllers

// swagger:parameters GetPersonByIDRequestParameter
type GetPersonByIDRequestParameter struct {
	// in:path
	// id person
	IdPerson int `json:"id_person"`

	// age
	Age int `json:"age"`
}

// swagger:parameters GetPersonAddressByIDRequestParameter
type GetPersonAddressByIDRequestParameter struct {
	// in:path
	// id person
	IdPerson int `json:"id_person"`

	// in:path
	// id address
	IdAddress int `json:"id_address"`
}
