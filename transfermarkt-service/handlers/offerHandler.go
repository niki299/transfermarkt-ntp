package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"transfermarkt/services"
)

func FindOffers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	offers, err := services.FindOffers(request)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Ok")
		log.Printf("Sent")
	}
	log.Printf("Sent")

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(offers)
}
