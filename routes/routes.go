package routers

import (
	"net/http"

	"example.com/controllers"
)

// SetRoutes sets up routing for the API
func SetRoutes() {
  http.HandleFunc("/projects", controllers.CreateProject)
  http.HandleFunc("/project", controllers.GetProjects)
  
}
