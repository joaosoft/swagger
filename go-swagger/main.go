package main

import (
	"fmt"
	"github.com/joaosoft/swagger/go-swagger/routes"
	_ "github.com/joaosoft/swagger/go-swagger/routes"
	"net/http"
)

func main() {
	fmt.Println("server started at http://localhost:8081/swaggerui/")
	if err := http.ListenAndServe(":8081", routes.Router); err != nil {
		panic(err)
	}
}
