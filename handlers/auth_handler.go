package handlers

import "database/sql"

// AuthHandler regroupe tous les handlers d'authentification.
type AuthHandler struct {
	db *sql.DB
}

// NewAuthHandler crée un AuthHandler avec la connexion DB fournie.
func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}
