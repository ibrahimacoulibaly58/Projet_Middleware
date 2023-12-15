package users

import (
    "encoding/json"
    "net/http"

    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "Projet_Middleware/User/internal/models"
    "Projet_Middleware/User/internal/services/users"
)

// DeleteUser
// @Tags         users
// @Summary      Delete a user.
// @Description  Delete a user.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    userID, _ := ctx.Value("userID").(uuid.UUID)

    
    err := users.DeleteUser(userID)
    if err != nil {
        logrus.Errorf("error: %s", err.Error())

        switch err.(type) {
        case *models.CustomError:
            // Erreur personnalisée
            customError := err.(*models.CustomError)
            w.WriteHeader(customError.Code)
            body, _ := json.Marshal(customError)
            _, _ = w.Write(body)
        // case users.ErrUserNotFound:
            // Utilisateur non trouvé
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
