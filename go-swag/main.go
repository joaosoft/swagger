package main

import (
	"fmt"
	"net/http"
	"swagger/go-swagger/routes"
	_ "swagger/go-swagger/routes"
)

func main() {
	fmt.Println("server started at http://localhost:8081/swaggerui/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
