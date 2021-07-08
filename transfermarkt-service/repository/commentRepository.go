package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"transfermarkt/model"
)

func CreateComment(comment *model.Comment) (*model.Comment, error) {
	fmt.Println("----------- comment creating ---------")

	db, err := sql.Open("postgres", pgConnStr)
	check(err)

	insertQuery := `insert into "comment"("id", "offer_id", "content") 
					values($1, $2, $3)`

	_, er := db.Exec(insertQuery, comment.Id, comment.OfferId, comment.Content)
	check(er)

	return comment, nil
}
