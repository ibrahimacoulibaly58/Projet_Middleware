package models

import "github.com/gofrs/uuid"

type Song struct {
    ID     *uuid.UUID `json:"id"`
    Title  string     `json:"title"`
    Artist string     `json:"artist"`
    
}

// SongCreateRequest représente les données nécessaires pour créer une nouvelle chanson
type SongCreateRequest struct {
    Title  string
    Artist string
}

// SongUpdateRequest représente les données nécessaires pour mettre à jour une chanson existante
type SongUpdateRequest struct {
    Title  string
    Artist string
}