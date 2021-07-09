package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"transfermarkt/handlers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/players", handlers.GetAllPlayers)
	router.HandleFunc("/offers-search", handlers.FindOffers)
	//filter
	router.HandleFunc("/rate", handlers.SaveRate)
	router.HandleFunc("/comment", handlers.SaveComment)

	log.Printf("Listening on port 80")
	http.ListenAndServe(":80", router)
}
