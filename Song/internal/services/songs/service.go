package songs

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/Song/internal/models"
	repository "Projet_Middleware/Song/internal/repositories/songs"
	"net/http"
)

func GetAllSongs() ([]models.Song, error) {
	var err error
	// Appel du dépôt
	songs, err := repository.GetAllSongs()
	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("erreur lors de la récupération des chansons : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return songs, nil
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	song, err := repository.GetSongById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "chanson non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la récupération de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, nil
}

func CreateSong(request models.SongCreateRequest) (*models.Song, error) {
	// Appel du dépôt
	song, err := repository.CreateSong(&request)
	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("erreur lors de la création de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, nil
}

func UpdateSong(id uuid.UUID, request models.SongUpdateRequest) (*models.Song, error) {
	// Appel du dépôt
	song, err := repository.UpdateSong(id, &request)
	// Gestion des erreurs
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "chanson non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la mise à jour de la chanson : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return song, nil
}

func DeleteSong(id uuid.UUID) error {
	// Appel du dépôt
	err := repository.DeleteSong(id)
	// Gestion des erreurs
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "chanson non trouvée",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la suppression de la chanson : %s", err.Error())
		return &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}