package songs

import (
	"encoding/json"
	"net/http"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/Song/internal/models"
	"Projet_Middleware/Song/internal/services/songs" 
)

// GetSongs
// @Tags         songs
// @Summary      Get songs.
// @Description  Get songs.
// @Success      200            {array}  models.Song
// @Failure      500             "Something went wrong"
// @Router       /songs [get]
func GetSongs(w http.ResponseWriter, _ *http.Request) {
	// Appeler le service
	songs, err := songs.GetAllSongs() // Assurez-vous que la fonction GetAllSongs existe dans le service
	if err != nil {
		// Journaliser l'erreur
		logrus.Errorf("error: %s", err.Error())

		switch err.(type) {
		case *models.CustomError:
			// Erreur personnalisée
			customError := err.(*models.CustomError)
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		default:
			// Erreur interne du serveur
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(songs)
	_, _ = w.Write(body)
}