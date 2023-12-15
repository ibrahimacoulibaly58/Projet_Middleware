package models

import "github.com/gofrs/uuid"

type User struct {
    ID       *uuid.UUID `json:"id"`
    Username string     `json:"username"`
}

// UserCreateRequest représente les données nécessaires pour créer un nouvel utilisateur
type UserCreateRequest struct {
    Username string
}

// UserUpdateRequest représente les données nécessaires pour mettre à jour un utilisateur existant
type UserUpdateRequest struct {
    Username string
}
