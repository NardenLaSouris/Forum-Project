package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/database"
	"forum/handlers"
)

func main() {
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/register-page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./templates/register.html")
	})

	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	err = database.CreateTables()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
