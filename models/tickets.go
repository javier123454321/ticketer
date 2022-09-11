package models

import (
	"database/sql"
	"fmt"
)

type Ticket struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	DueDate     string  `json:"due_date"`
	Description *string `json:"line_items"`
	Title       string  `jsont:"title"`
	Status      string  `json:"status"`
}

func (t *Ticket) CreateTable(db *sql.DB) {
	stmnt1 := `CREATE TYPE ticket_states AS ENUM ('in backlog', 'to-do', 'doing', 'done');`
	stmt, err := db.Prepare(stmnt1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	fmt.Println("creating ticket_states enum")
	_, err = db.Exec(stmnt1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stmnt := `CREATE TABLE "tickets" (
		id              serial primary key,
		due_date        timestamp,
		user_id         integer references "users"(id),
		description		text,
		title			varchar(255) NOT NULL,
		status			ticket_states DEFAULT 'in backlog'
);`
	stmt, err = db.Prepare(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	fmt.Println("creating tickets table")
	_, err = db.Exec(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success!")
}

func (t *Ticket) Retrieve(db *sql.DB, id int) (err error) {
	err = db.QueryRow("select id, user_id, due_date, status, title, description from tickets where id = $1", id).Scan(&t.Id, &t.UserId, &t.DueDate, &t.Status, &t.Title, &t.Description)
	return err
}

func (t *Ticket) Create(db *sql.DB) (err error) {
	statement := `insert into "tickets" (due_date, user_id, title, description) values ($1, $2, $3, $4) returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(t.DueDate, t.UserId, t.Title, t.Description).Scan(&t.Id)
	return err
}

func (t *Ticket) Index(db *sql.DB, limit int) ([]Ticket, error) {
	tickets := []Ticket{}
	rows, err := db.Query("select id, due_date, user_id, title, status from tickets limit $1;",
		limit)
	if err != nil {
		return []Ticket{}, err
	}

	defer rows.Close()

	for rows.Next() {
		ticket := Ticket{}
		err = rows.Scan(&ticket.Id, &ticket.DueDate, &ticket.UserId, &ticket.Title, &ticket.Status)
		if err != nil {
			return []Ticket{}, err
		}
		fmt.Println("checking", ticket)
		tickets = append(tickets, ticket)
	}
	return tickets, nil
}
