package handlers

import (
	"encoding/json"
	"net/http"
	"transfermarkt/services"
)

func SaveRate(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	rate, err := services.CreateRate(request)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Saving rate error")
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(rate)
}
