package controllers

import (
	"fmt"
	"github.com/labstack/echo"
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
func GetPersonAddressByID(ctx echo.Context) error {
	request := GetPersonAddressByIDRequest{
		IdPerson:  ctx.Param(":id_person"),
		IdAddress: ctx.Param(":id_address"),
	}

	fmt.Printf("> executing get address for id_person: %s, id_address: %s", request.IdPerson, request.IdAddress)

	response :=	AddressResponse{
			Id:      request.IdAddress,
			Country: "Portugal",
			City:    "Porto",
			Street:  "Rua da calçada",
			Number:  7,
		}

	return ctx.JSON(http.StatusOK, response)
}
