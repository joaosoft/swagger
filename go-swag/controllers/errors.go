package controllers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// GetErrorByID example
// @Summary Returns the error definition
// @Description Returns the error definition
// @ID get-error-by-id
// @Accept  json
// @Produce  json
// @Param   id_error      query   int     true  "Error ID"
// @Success 200 {object} ErrorResponse "ok"
// @Success 204 {string} string	"ok"
// @Router /errors [get]
func GetErrorByID(ctx echo.Context) error {
	errorID, _ := strconv.Atoi(ctx.QueryParam("id_error"))
	fmt.Printf("> executing get errors for id: %d", errorID)

	statusText := http.StatusText(errorID)

	if statusText != "" {
		response := ErrorResponse{
				Code:    errorID,
				Message: statusText,
			}
		return ctx.JSON(http.StatusOK, response)
	} else {
		return ctx.NoContent(http.StatusNoContent)
	}
}
