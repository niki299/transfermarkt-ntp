package handlers

import (
	"encoding/json"
	"net/http"
	"transfermarkt/services"
)

func SaveComment(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	comment, err := services.CreateComment(request)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Saving comment error")
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(comment)
}
