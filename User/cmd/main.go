package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/User/internal/controllers/users" // Importez le package des contrôleurs d'utilisateurs
	"Projet_Middleware/User/internal/helpers"
	_"Projet_Middleware/User/internal/models" 
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// GET /users - Get all users
	r.Get("/users", users.GetUsers)

	// POST /users - Create a new user
	r.Post("/users", users.CreateUser)

	// GET /users/{id} - Get a specific user by ID
	r.Get("/users/{id}", users.GetUser)

	// PUT /users/{id} - Update a specific user by ID
	r.Put("/users/{id}", users.UpdateUser)

	// DELETE /users/{id} - Delete a specific user by ID
	r.Delete("/users/{id}", users.DeleteUser)

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			username VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
