package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// swagger:route GET /persons/{id_person}/addresses/{id_address} address-tag GetPersonAddressByIDRequestParameter
//
// Get the person address.
//
// This will return error.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http, https
//
//     Deprecated: false
//
//     Responses:
//       default: AddressResponse
//       200: AddressResponse
//       400: ErrorResponse
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
