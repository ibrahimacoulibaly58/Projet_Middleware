package songs

import (
    "encoding/json"
    "net/http"

    "github.com/sirupsen/logrus"
    "Projet_Middleware/Song/internal/models"
    "Projet_Middleware/Song/internal/services/songs"
)

// CreateSong
// @Tags         songs
// @Summary      Create a new song.
// @Description  Create a new song.
// @Param        request       body      models.SongCreateRequest  true  "Song creation request"
// @Success      201            {object}  models.Song
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /songs [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
    var createRequest models.SongCreateRequest
    if err := json.NewDecoder(r.Body).Decode(&createRequest); err != nil {
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

    
    song, err := songs.CreateSong(createRequest)
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

    // Répondre avec la chanson créée
    w.WriteHeader(http.StatusCreated)
    body, _ := json.Marshal(song)
    _, _ = w.Write(body)
}