package main

import (
	"log"
	"net/http"

	"forum/database"
	"forum/handlers"
)

http.HandleFunc("/register", handlers.RegisterHandler)

http.HandleFunc("/register-page", func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./templates/register.html")
})

func main() {

	err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/register", handlers.RegisterHandler)

	fmt.Println("Server running on :8080")

	http.ListenAndServe(":8080", nil)
}