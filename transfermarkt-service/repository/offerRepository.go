package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"time"
	"transfermarkt/model"
)

func FindAll(searchData *model.SearchData) ([]model.OfferResponse, error) {

	db, err := sql.Open("postgres", pgConnStr)
	check(err)

	defer db.Close()

	rows, err := db.Query(defaultQuery, searchData.Name, searchData.Surname, searchData.Club,
		searchData.Position, searchData.PriceFrom, searchData.PriceTo)
	check(err)
	defer rows.Close()

	var offers []model.OfferResponse

	for rows.Next() {
		var offer_id string
		var date time.Time
		var name string
		var surname string
		var age int
		var foot string
		var club string
		var value int
		var position string
		var avg_rate float32
		var comments []model.Comment

		err := rows.Scan(&offer_id, &date, &name, &surname, &age, &foot, &club, &value, &position)
		check(err)

		row_rate, err := db.Query(averageRate, offer_id)
		check(err)
		// defer row_rate.Close()

		var some_id string
		if row_rate.Next() {
			rate_err := row_rate.Scan(&some_id, &avg_rate)
			check(rate_err)
		}
		log.Print(avg_rate)

		row_comment, err := db.Query(commentQuery, offer_id)
		check(err)
		for row_comment.Next() {
			var content string
			comment_err := row_comment.Scan(&content)
			check(comment_err)
			comments = append(comments, model.Comment{OfferId: some_id, Content: content})
		}

		offers = append(offers, model.OfferResponse{Id: offer_id,
			Player: model.Player{Name: name, Surname: surname, Age: age,
				Foot: foot, Club: club, PlayerPosition: model.Position(position), PlayerValue: value},
			Date: date, AverageRate: avg_rate, Comments: comments})
	}
	return offers, nil
}

func FindAllOffers() ([]model.Offer, error) {

	return nil, nil
}
