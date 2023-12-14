package songs

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/Song/internal/models"
	"Projet_Middleware/Song/internal/repositories/songs"
)

// GetSong
// @Tags         songs
// @Summary      Get a song.
// @Description  Get a song.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Success      200            {object}  models.Song
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [get]
func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	songID, _ := ctx.Value("songID").(uuid.UUID)

	// Assurez-vous que la fonction GetSongByID existe dans le référentiel
	song, err := songs.GetSongById(songID)
	if err != nil {
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

	// Répondre avec la chanson récupérée
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
}