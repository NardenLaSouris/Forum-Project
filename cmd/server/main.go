package main

import (
	"log"
	"net/http"
	"time"

	"Forum-Project/config"
	"Forum-Project/database"
	"Forum-Project/routes"
	"Forum-Project/sessions"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.InitDB(cfg.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	database.DB = db

	sessions.Init()

	router := routes.SetupRoutes()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
