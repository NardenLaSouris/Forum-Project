package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"forum/database"
	"forum/utils"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	// Vérifie méthode POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Récupération données formulaire
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Vérification champs vides
	if username == "" || email == "" || password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Vérifie si email déjà utilisé
	var existingID int

	err := database.DB.QueryRow(
		"SELECT id FROM users WHERE email = ?",
		email,
	).Scan(&existingID)

	// SI email trouvé
	if err == nil {
		http.Error(w, "Email already used", http.StatusConflict)
		return
	}

	// SI autre erreur SQL
	if err != sql.ErrNoRows {
		http.Error(w, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Hash du mot de passe
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insertion utilisateur
	_, err = database.DB.Exec(
		`
		INSERT INTO users(username, email, password)
		VALUES (?, ?, ?)
		`,
		username,
		email,
		hashedPassword,
	)

	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	// Succès
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
