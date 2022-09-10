package dbconfig

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	var err error
	db, err := sql.Open("postgres", "dbname=ticketer sslmode=disable")
	if err != nil {
		fmt.Println((err.Error()))
		panic(err)
	}
	return db
}
