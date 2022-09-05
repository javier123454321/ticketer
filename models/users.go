package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (u *User) CreateTable(db *sql.DB) {
	stmnt := `CREATE TABLE "user" (
		id              serial primary key,
		email           varchar(255),
		name            varchar(255)
);`
	stmt, err := db.Prepare(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	fmt.Println("creating users table")
	_, err = db.Exec(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success!")
}

func (u *User) Create(db *sql.DB) error {
	stmnt := `insert into "user" (name, email) values ($1, $2) returning id`

	stmt, err := db.Prepare(stmnt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.Name, u.Email).Scan(&u.Id)
	return err
}
