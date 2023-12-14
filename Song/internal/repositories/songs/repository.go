package songs

import (
	
	"github.com/gofrs/uuid"
	"Projet_Middleware/Song/internal/helpers"
	"Projet_Middleware/Song/internal/models"
)

// GetAllSongs récupère toutes les chansons
func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM songs")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// Parsing des données dans une slice d'objets Song
	songs := []models.Song{}
	for rows.Next() {
		var song models.Song
		err = rows.Scan(&song.ID, &song.Title, &song.Artist)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	// N'oubliez pas de fermer les lignes
	_ = rows.Close()

	return songs, err
}

// GetSongById récupère une chanson par ID
func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM songs WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.ID, &song.Title, &song.Artist)
	if err != nil {
		return nil, err
	}
	return &song, err
}

// CreateSong crée une nouvelle chanson
func CreateSong(songRequest *models.SongCreateRequest) (*models.Song, error) {
	// Générer une nouvelle ID pour la chanson
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Créer une instance de Song avec les données de la requête
	newSong := &models.Song{
		ID:     &id,
		Title:  songRequest.Title,
		Artist: songRequest.Artist,
	}

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	// Insérer la nouvelle chanson dans la base de données
	_, err = db.Exec("INSERT INTO songs (id, title, artist) VALUES (?, ?, ?)",
		newSong.ID.String(), newSong.Title, newSong.Artist)
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	return newSong, nil
}

// UpdateSong met à jour une chanson existante
func UpdateSong(id uuid.UUID, songRequest *models.SongUpdateRequest) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	// Mettre à jour la chanson dans la base de données
	_, err = db.Exec("UPDATE songs SET title = ?, artist = ? WHERE id = ?",
		songRequest.Title, songRequest.Artist, id.String())
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	// Récupérer la chanson mise à jour
	updatedSong, err := GetSongById(id)
	if err != nil {
		return nil, err
	}

	return updatedSong, nil
}

// DeleteSong supprime une chanson par ID
func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	// Supprimer la chanson de la base de données
	_, err = db.Exec("DELETE FROM songs WHERE id = ?", id.String())
	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return nil
}