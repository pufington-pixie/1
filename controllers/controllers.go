package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/database"
	"example.com/models"
	"github.com/gorilla/mux"
)

// POST /projects
func CreateProject(w http.ResponseWriter, r *http.Request) {
  // parse request body and validate
  var project models.Project
  err := json.NewDecoder(r.Body).Decode(&project)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  // create new Project model
  database.CreateProject(project)
  // return 201 Created
  w.WriteHeader(http.StatusCreated)
}

// GET /projects
func GetProjects(w http.ResponseWriter, r *http.Request) {
  // query database for list of projects
  projects := database.GetProjects()
  // return projects as JSON
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(projects)
}

// GET /projects/{id}
func GetProject(w http.ResponseWriter, r *http.Request) {
  // get project ID from URL params
  id := getProjectID(r)
  // query database and get project with that ID
  project := database.GetProject(id)
  if project == nil {
    // return 404 if not found
    w.WriteHeader(http.StatusNotFound)
    return
  }
  // return project as JSON 
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(project)
}

// PUT /projects/{id}
func UpdateProject(w http.ResponseWriter, r *http.Request) {
  // get project ID from URL params
  id := getProjectID(r)
  // parse request body and validate
  var project models.Project
  err := json.NewDecoder(r.Body).Decode(&project)
  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    return
  }
  // query database and get project with that ID
  projectPtr := database.GetProject(id)
  if projectPtr == nil {
    // return 404 if not found
    w.WriteHeader(http.StatusNotFound)
    return
  }
  // update project fields
  projectPtr.Name = project.Name   // etc...
  // save updated project to database
  database.UpdateProject(*projectPtr)
  // return 204 No Content
  w.WriteHeader(http.StatusNoContent)
}

// DELETE /projects/{id}
func DeleteProject(w http.ResponseWriter, r *http.Request) {
  // get project ID from URL params
  id := getProjectID(r)
  // query database and get project with that ID
  project := database.GetProject(id)
  if project == nil {
    // return 404 if not found
    w.WriteHeader(http.StatusNotFound)
    return
  }
  // delete project from database
  database.DeleteProject(id)
  // return 204 No Content
  w.WriteHeader(http.StatusNoContent) 
}

func getProjectID(r *http.Request) int {
  // get project ID from URL params
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  return id
}
