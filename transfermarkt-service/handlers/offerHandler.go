package handlers

import (
	"encoding/json"
	"net/http"
	"transfermarkt/services"
)

func FindOffers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	offers, err := services.FindOffers(request)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Ok")
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(offers)
}
