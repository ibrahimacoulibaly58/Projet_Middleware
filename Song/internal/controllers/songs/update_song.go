package songs

import (
    "encoding/json"
    "net/http"

    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "Projet_Middleware/Song/internal/models"
    "Projet_Middleware/Song/internal/services/songs"
)

// UpdateSong
// @Tags         songs
// @Summary      Update a song.
// @Description  Update a song.
// @Param        id            path      string  true  "Song UUID formatted ID"
// @Param        request       body      models.SongUpdateRequest  true  "Song update request"
// @Success      200            {object}  models.Song
// @Failure      400            "Invalid request payload"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    songID, _ := ctx.Value("songID").(uuid.UUID)

    var updateRequest models.SongUpdateRequest
    if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
        logrus.Errorf("parsing error: %s", err.Error())
        customError := &models.CustomError{
            Message: "invalid request payload",
            Code:    http.StatusBadRequest,
        }
        w.WriteHeader(customError.Code)
        body, _ := json.Marshal(customError)
        _, _ = w.Write(body)
        return
    }

    // Assurez-vous que la fonction UpdateSong existe dans le service
    song, err := songs.UpdateSong(songID, updateRequest)
    if err != nil {
        logrus.Errorf("error: %s", err.Error())

        switch err.(type) {
        case *models.CustomError:
            // Erreur personnalisée
            customError := err.(*models.CustomError)
            w.WriteHeader(customError.Code)
            body, _ := json.Marshal(customError)
            _, _ = w.Write(body)
        //case songs.ErrSongNotFound:
            // Chanson non trouvée
           // w.WriteHeader(http.StatusNotFound)
        default:
            // Erreur interne du serveur
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }

    // Répondre avec la chanson mise à jour
    w.WriteHeader(http.StatusOK)
    body, _ := json.Marshal(song)
    _, _ = w.Write(body)
}