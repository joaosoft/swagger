package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// GetPersonAddressByIDRequestParameter example
// @Summary Get the person address.
// @Description Get the person address.
// @ID get-person-address-by-id
// @Accept  json
// @Produce  json
// @Param   id_person      path   int     true  "Person ID"
// @Param   id_address      path   int     true  "Person ID"
// @Success 200 {string} AddressResponse	"ok"
// @Failure 400 {object} ErrorResponse "error"
// @Router /persons/{id_person}/addresses/{id_address} [get]
func GetPersonAddressByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", vars["id_person"], vars["id_address"])

	bytes, _ := json.Marshal(
		AddressResponse{
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
