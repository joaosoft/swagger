package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GetPersonAddressByIDRequest example
// @Summary Get the person address.
// @Description Get the person address.
// @ID get-person-address-by-id
// @Accept  json
// @Produce  json
// @Param   id_person      path   string     true  "Person ID"
// @Param   id_address      path   string     true  "Address ID"
// @Success 200 {string} AddressResponse	"ok"
// @Failure 400 {object} ErrorResponse "error"
// @Router /persons/{id_person}/addresses/{id_address} [get]
func GetPersonAddressByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	request := GetPersonAddressByIDRequest{
		IdPerson:  vars["id_person"],
		IdAddress: vars["id_address"],
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	bytes, _ := json.Marshal(
		AddressResponse{
			Id:      request.IdAddress,
			Country: "Portugal",
			City:    "Porto",
			Street:  "Rua da calçada",
			Number:  7,
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
