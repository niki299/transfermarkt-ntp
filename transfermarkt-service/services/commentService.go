package services

import (
	"log"
	"net/http"
	"transfermarkt/model"
	"transfermarkt/repository"

	"github.com/google/uuid"
)

func CreateComment(request *http.Request) (*model.Comment, error) {
	commentReq := model.Comment{}
	err := parse(request, &commentReq)
	if err != nil {
		log.Printf("Cannot parse rate body.")
		return nil, err
	}
	commentReq.Id = uuid.New().String()

	return repository.CreateComment(&commentReq)
}
