package services

import (
	"log"
	"net/http"
	"transfermarkt/model"
	"transfermarkt/repository"
)

func FindOffers(request *http.Request) ([]model.OfferResponse, error) {
	searchData := model.SearchData{}

	err := parse(request, &searchData)
	if err != nil {
		log.Printf("Cannot parse search body.")
		return nil, err
	}

	return repository.FindAll(&searchData)
}
