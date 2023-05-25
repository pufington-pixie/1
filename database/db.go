package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
 
func ConnectDB() (*sql.DB,error) {
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := "ganesh"
    dbName := "project"
	
    // Open the database connection
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+ "?parseTime=true")
    if err != nil {
        log.Fatal(err)
    }
   

    // Ping the database to ensure the connection is established
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // Connection established
    fmt.Println("Connected to the database!")
    return db,nil
    
}
// package database

// import (
// 	"database/sql"
// 	_ "github.com/go-sql-driver/mysql"

// )

// // DB is the database connection
// var DB *sql.DB

// // Connect opens a database connection
// func Connect() error {
// 	dsn := "root:ganesh@tcp(127.0.0.1:3306)/project?parseTime=true"
// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 	  return err
// 	}
// 	DB = db
// 	return nil
//   }
// // // CreateProject inserts a new project into the database
// // func CreateProject(project models.Project) {
// //   stmt, err := DB.Prepare("INSERT INTO projects (name, title, date, sap_number, notes, branch_id, status_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// //   _, err = stmt.Exec(project.Name, project.Title, project.Date, project.SAPNumber, project.Notes, project.BranchID, project.StatusID)
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// // }

// // // GetProjects returns a list of all projects
// // func GetProjects() []models.Project {
// //   rows, err := DB.Query("SELECT * FROM projects")
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// //   projects := []models.Project{}
// //   for rows.Next() {
// //     project := models.Project{}
// //     err = rows.Scan(&project.ID, &project.Name, &project.Title, &project.Date, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID)
// //     if err != nil {
// //       log.Fatal(err)
// //     }
// //     projects = append(projects, project)
// //   }
// //   return projects
// // }  

// // // GetProject returns a project by ID
// // func GetProject(id int) *models.Project {
// //   project := models.Project{}
// //   err := DB.QueryRow("SELECT * FROM projects WHERE id= ? ", id).Scan(&project.ID, &project.Name, &project.Title, &project.Date, &project.SAPNumber, &project.Notes, &project.BranchID, &project.StatusID)
// //   if err != nil {
// //     return nil
// //   }
// //   return &project
// // }

// // // UpdateProject updates a project in the database
// // func UpdateProject(project models.Project) {
// //   stmt, err := DB.Prepare("UPDATE projects SET name= ? title=?, date=?, sap_number=?, notes=?, branch_id=?, status_id=? WHERE id=?")
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// //   _, err = stmt.Exec(project.Name, project.Title, project.Date, project.SAPNumber, project.Notes, project.BranchID, project.StatusID, project.ID)
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// // }

// // // DeleteProject deletes a project from the database
// // func DeleteProject(id int) {
// //   stmt, err := DB.Prepare("DELETE FROM projects WHERE id=?")
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// //   _, err = stmt.Exec(id)
// //   if err != nil {
// //     log.Fatal(err)
// //   }
// // }
