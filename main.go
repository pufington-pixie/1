package main

import (
	"log"
	"net/http"

	"example.com/database"
	"example.com/routes"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	routers.SetRoutes()
	http.ListenAndServe(":8000", nil)
}
