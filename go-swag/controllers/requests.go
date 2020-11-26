package controllers

// GetPersonByIDRequestParameter
type GetPersonByIDRequestParameter struct {
	// in:path
	// id person
	IdPerson int `json:"id_person"`

	// age
	Age int `json:"age"`
}

// GetPersonAddressByIDRequestParameter
type GetPersonAddressByIDRequestParameter struct {
	// in:path
	// id person
	IdPerson int `json:"id_person"`

	// in:path
	// id address
	IdAddress int `json:"id_address"`
}