package users

import (
	"encoding/json"
	"net/http"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/User/internal/models"
	"Projet_Middleware/User/internal/services/users"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /users [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// Appeler le service
	users, err := users.GetAllUsers() 
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
	body, _ := json.Marshal(users)
	_, _ = w.Write(body)
}
