package users

import (
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/User/internal/models"
	"Projet_Middleware/User/internal/repositories/users"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userID, _ := ctx.Value("userID").(uuid.UUID)

	
	user, err := users.GetUserById(userID)
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

	// Répondre avec l'utilisateur récupéré
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(user)
	_, _ = w.Write(body)
}
