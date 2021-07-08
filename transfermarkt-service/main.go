package main

import (
	"log"
	"net/http"
	"transfermarkt/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players", handlers.GetAllPlayers)
	router.HandleFunc("/rate", handlers.SaveRate)
	router.HandleFunc("/comment", handlers.SaveComment)

	log.Printf("Listening on port 80")
	http.ListenAndServe(":80", router)
}
