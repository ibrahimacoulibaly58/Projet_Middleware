package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/Song/internal/controllers/songs" 
	"Projet_Middleware/Song/internal/helpers"
	_"Projet_Middleware/Song/internal/models" 
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// GET /songs - Get all songs
	r.Get("/songs", songs.GetSongs)

	// POST /songs - Create a new song
	r.Post("/songs", songs.CreateSong)

	// GET /songs/{id} - Get a specific song by ID
	r.Get("/songs/{id}", songs.GetSong)

	// PUT /songs/{id} - Update a specific song by ID
	r.Put("/songs/{id}", songs.UpdateSong)

	// DELETE /songs/{id} - Delete a specific song by ID
	r.Delete("/songs/{id}", songs.DeleteSong)

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8080", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			title VARCHAR(255) NOT NULL,
			artist VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}