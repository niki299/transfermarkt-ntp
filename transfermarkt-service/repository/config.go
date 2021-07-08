package repository

import "fmt"

var (
	dbUsername = "postgres"
	dbPassword = "postgres"
	dbTable    = "postgres"
	dbHost     = "localHost"
	dbPort     = "5438"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
)
