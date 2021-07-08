package handlers

import (
	"encoding/json"
	"net/http"
	"transfermarkt/repository"
)

func GetAllPlayers(response http.ResponseWriter, request *http.Request) {
	// vars := mux.Vars(request)
	// title := vars["title"]
	// page := vars["page"]

	response.Header().Set("Content-Type", "application/json")

	players, err := repository.FindAllPlayers()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode("Error getting the players")
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(players)
}
