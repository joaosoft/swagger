package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, req *http.Request) {
	fmt.Println("> executing not found")

	bytes, _ := json.Marshal(
		ErrorResponse{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		},
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(bytes)
}
