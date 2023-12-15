package users

import (
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"Projet_Middleware/User/internal/models"
	repository "Projet_Middleware/User/internal/repositories/users"
	"net/http"
)

func GetAllUsers() ([]models.User, error) {
	var err error
	// Appel du dépôt
	users, err := repository.GetAllUsers()
	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("erreur lors de la récupération des utilisateurs : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return users, nil
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "utilisateur non trouvé",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la récupération de l'utilisateur : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func CreateUser(request models.UserCreateRequest) (*models.User, error) {
	// Appel du dépôt
	user, err := repository.CreateUser(&request)
	// Gestion des erreurs
	if err != nil {
		logrus.Errorf("erreur lors de la création de l'utilisateur : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func UpdateUser(id uuid.UUID, request models.UserUpdateRequest) (*models.User, error) {
	// Appel du dépôt
	user, err := repository.UpdateUser(id, &request)
	// Gestion des erreurs
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "utilisateur non trouvé",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la mise à jour de l'utilisateur : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return user, nil
}

func DeleteUser(id uuid.UUID) error {
	// Appel du dépôt
	err := repository.DeleteUser(id)
	// Gestion des erreurs
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "utilisateur non trouvé",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("erreur lors de la suppression de l'utilisateur : %s", err.Error())
		return &models.CustomError{
			Message: "Quelque chose s'est mal passé",
			Code:    http.StatusInternalServerError,
		}
	}

	return nil
}
