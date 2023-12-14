package songs

import (
    "encoding/json"
    "net/http"

    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "Projet_Middleware/Song/internal/models"
    "Projet_Middleware/Song/internal/services/songs"
)

// DeleteSong
// @Tags         songs
// @Summary      Delete a song.
// @Description  Delete a song.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    songID, _ := ctx.Value("songID").(uuid.UUID)

    // Assurez-vous que la fonction DeleteSong existe dans le service
    err := songs.DeleteSong(songID)
    if err != nil {
        logrus.Errorf("error: %s", err.Error())

        switch err.(type) {
        case *models.CustomError:
            // Erreur personnalisée
            customError := err.(*models.CustomError)
            w.WriteHeader(customError.Code)
            body, _ := json.Marshal(customError)
            _, _ = w.Write(body)
       // case songs.ErrSongNotFound:
            // Chanson non trouvée
            //w.WriteHeader(http.StatusNotFound)
        default:
            // Erreur interne du serveur
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    // Répondre avec succès (pas de contenu)
    w.WriteHeader(http.StatusNoContent)
}