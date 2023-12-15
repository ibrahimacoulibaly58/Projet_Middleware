package users

import (
    "encoding/json"
    "net/http"
    "github.com/gofrs/uuid"
    "github.com/sirupsen/logrus"
    "Projet_Middleware/User/internal/models"
    "Projet_Middleware/User/internal/services/users"
)

// UpdateUser
// @Tags         users
// @Summary      Update a user.
// @Description  Update a user.
// @Param        id            path      string  true  "User UUID formatted ID"
// @Param        request       body      models.UserUpdateRequest  true  "User update request"
// @Success      200            {object}  models.User
// @Failure      400            "Invalid request payload"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    userID, _ := ctx.Value("userID").(uuid.UUID)

    var updateRequest models.UserUpdateRequest
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

    user, err := users.UpdateUser(userID, updateRequest)
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

    // Répondre avec l'utilisateur mis à jour
    w.WriteHeader(http.StatusOK)
    body, _ := json.Marshal(user)
    _, _ = w.Write(body)
}
