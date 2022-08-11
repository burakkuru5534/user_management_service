package helper

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "tayitkan"
	dbname   = "rollic"
)

func ConnectDb() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	// close database
	defer db.Close()

	// check db
	err = db.Ping()

	fmt.Println("Connected!")

	return db
}
