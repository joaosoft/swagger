package main

import (
	"fmt"
	"net/http"
	_ "swagger/routes"
)

func main() {
	fmt.Println("server started at http://localhost:8081/swaggerui/")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
