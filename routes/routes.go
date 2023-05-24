package routers

import (
	"net/http"

	"example.com/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// SetRoutes sets up the routing for the API
func SetRoutes() {
  r := chi.NewRouter()
  
  // Middleware
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  
  // Routes
  r.Post("/project", controllers.CreateProject)
  r.Get("/projects", controllers.GetProjects)
  r.Get("/projects/{id}", controllers.GetProject)
  r.Put("/projects/{id}", controllers.UpdateProject)
  r.Delete("/projects/{id}", controllers.DeleteProject)
  
  // Serve
  http.ListenAndServe(":8000", r)
}
