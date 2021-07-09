package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"transfermarkt/model"
)

func FindAllPlayers() ([]model.Player, error) {
	//open database
	db, err := sql.Open("postgres", pgConnStr)
	check(err)

	defer db.Close()

	rows, err := db.Query(`SELECT "player_id", "name", "surname", 
								  "age", "foot", "club", "playerposition", playervalue FROM "player"`)
	check(err)

	defer rows.Close()

	var players []model.Player

	for rows.Next() {
		var id string
		var name string
		var surname string
		var age int
		var foot string
		var club string
		var playerPosition model.Position
		var playerValue int

		err := rows.Scan(&id, &name, &surname, &age, &foot, &club, &playerPosition, &playerValue)
		check(err)

		players = append(players, model.Player{Id: id, Name: name, Surname: surname, Age: age,
			Foot: foot, Club: club, PlayerPosition: playerPosition, PlayerValue: playerValue})
	}

	return players, nil
}

func check(err error) {
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
}
