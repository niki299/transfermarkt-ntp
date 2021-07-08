package services

import (
	"encoding/json"
	"log"
	"net/http"
	"transfermarkt/model"
	"transfermarkt/repository"

	"github.com/google/uuid"
)

func CreateRate(request *http.Request) (*model.Rate, error) {
	rateReq := model.Rate{}
	err := parse(request, &rateReq)
	if err != nil {
		log.Printf("Cannot parse rate body")
		return nil, err
	}
	rateReq.Id = uuid.New().String()

	return repository.CreateRate(&rateReq)
}

func ckeck(err error) {
	if err != nil {
		panic(err)
	}
}

func parse(request *http.Request, data interface{}) error {
	return json.NewDecoder(request.Body).Decode(data)
}
