package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"transfermarkt/model"
)

func FindAll(searchData *model.SearchData) ([]model.Offer, error) {

	db, err := sql.Open("postgres", pgConnStr)
	check(err)

	defer db.Close()

	rows, err := db.Query(defaultQuery, searchData.Name, searchData.Surname, searchData.Club)
	check(err)
	defer rows.Close()

	var offers []model.Offer

	for rows.Next() {
		var offer_id string
		var date string
		var name string
		var surname string
		var age int
		var foot string
		var club string

		err := rows.Scan(&offer_id, &date, &name, &surname, &age, &foot, &club)
		check(err)

		offers = append(offers, model.Offer{Id: offer_id,
			Player: model.Player{Name: name, Surname: surname, Age: age,
				Foot: foot, Club: club}})
	}
	return offers, nil
}
