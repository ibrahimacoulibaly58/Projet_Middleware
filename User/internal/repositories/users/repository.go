package users

import (
	"github.com/gofrs/uuid"
	"Projet_Middleware/User/internal/helpers"
	"Projet_Middleware/User/internal/models"
)

// GetAllUsers récupère tous les utilisateurs
func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// Parsing des données dans une slice d'objets User
	users := []models.User{}
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// N'oubliez pas de fermer les lignes
	_ = rows.Close()

	return users, err
}

// GetUserById récupère un utilisateur par ID
func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var user models.User
	err = row.Scan(&user.ID, &user.Username)
	if err != nil {
		return nil, err
	}
	return &user, err
}

// CreateUser crée un nouvel utilisateur
func CreateUser(userRequest *models.UserCreateRequest) (*models.User, error) {
	// Générer une nouvelle ID pour l'utilisateur
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// Créer une instance de User avec les données de la requête
	newUser := &models.User{
		ID:       &id,
		Username: userRequest.Username,
	}

	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	// Insérer le nouvel utilisateur dans la base de données
	_, err = db.Exec("INSERT INTO users (id, username) VALUES (?, ?)",
		newUser.ID.String(), newUser.Username)
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}

// UpdateUser met à jour un utilisateur existant
func UpdateUser(id uuid.UUID, userRequest *models.UserUpdateRequest) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}

	// Mettre à jour l'utilisateur dans la base de données
	_, err = db.Exec("UPDATE users SET username = ? WHERE id = ?",
		userRequest.Username, id.String())
	helpers.CloseDB(db)

	if err != nil {
		return nil, err
	}

	// Récupérer l'utilisateur mis à jour
	updatedUser, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser supprime un utilisateur par ID
func DeleteUser(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	// Supprimer l'utilisateur de la base de données
	_, err = db.Exec("DELETE FROM users WHERE id = ?", id.String())
	helpers.CloseDB(db)

	if err != nil {
		return err
	}

	return nil
}
