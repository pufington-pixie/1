package main

import (
	"log"
	

	"example.com/database"
	"example.com/routes"
)

func main() {
	// Connect to database
	err := database.Connect()
	if err != nil {
	  log.Fatal(err)
	}
	
	// Set up routes
	routers.SetRoutes() 
  }
