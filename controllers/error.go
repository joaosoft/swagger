package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleError(w http.ResponseWriter, req *http.Request) {
	fmt.Println("> executing error")
	bytes, _ := json.Marshal(
		ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "error response",
		},
	)

	fmt.Fprint(w, string(bytes))
}
