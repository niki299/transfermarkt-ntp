package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"transfermarkt/model"
)

func CreateRate(rate *model.Rate) (*model.Rate, error) {
	fmt.Println("----------- rate creating ---------")

	db, err := sql.Open("postgres", pgConnStr)
	check(err)

	defer db.Close()

	insertQuery := `insert into "rate"("rate_id", "offer_id", "ratevalue") values($1, $2, $3)`

	_, er := db.Exec(insertQuery, rate.Id, rate.OfferId, rate.RateValue)
	check(er)

	return rate, nil
}

func ckeck(err error) {
	if err != nil {
		panic(err)
	}
}
