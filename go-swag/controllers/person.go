package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// GetPersonByIDRequest example
// @Summary Get person by id.
// @Description Get person by id.
// @ID get-person-by-id
// @Accept  json
// @Produce  json
// @Param   id_person      path   string     true  "Person ID"
// @Param   age      query   int     false  "Age"
// @Success 200 {string} PersonResponse	"ok"
// @Failure 400 {object} ErrorResponse "error"
// @Router /persons/{id_person} [get]
func GetPersonByID(ctx echo.Context) error {
	age, _ := strconv.Atoi(ctx.QueryParam("age"))
	request := GetPersonByIDRequest{
		IdPerson: ctx.Param(":id_person"),
		Age:      age,
	}

	fmt.Printf("> executing get person for id_person: %s", request.IdPerson)

	response := PersonResponse{
		Id:   request.IdPerson,
		Name: "João Ribeiro",
		Age:  request.Age,
	}

	return ctx.JSON(http.StatusOK, response)
}
