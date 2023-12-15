package users

import (
    "encoding/json"
    "net/http"

    "github.com/sirupsen/logrus"
    "Projet_Middleware/User/internal/models"
    "Projet_Middleware/User/internal/services/users"
)

// CreateUser
// @Tags         users
// @Summary      Create a new user.
// @Description  Create a new user.
// @Param        request       body      models.UserCreateRequest  true  "User creation request"
// @Success      201            {object}  models.User
// @Failure      400            "Invalid request payload"
// @Failure      500            "Something went wrong"
// @Router       /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var createRequest models.UserCreateRequest
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

    user, err := users.CreateUser(createRequest)
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

    // Répondre avec l'utilisateur créé
    w.WriteHeader(http.StatusCreated)
    body, _ := json.Marshal(user)
    _, _ = w.Write(body)
}
