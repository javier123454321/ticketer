package models

import (
	"database/sql"
	"fmt"
)

type Ticket struct {
	Id          uint   `json:"id"`
	PayerId     uint   `json:"payer_id"`
	DueDate     string `json:"due_date"`
	Description string `json:"line_items"`
	Status      string `json:"status"`
}

func (i *Ticket) CreateTable(db *sql.DB) {
	stmnt := `
CREATE TYPE ticket_states AS ENUM ('backlog', 'to-do', 'doing', 'done');
CREATE TABLE "tickets" (
		id              serial primary key,
		due_date        timestamp,
		user_id         integer references "user"(id),
		description		text,
		status			ticket_states DEFAULT 'backlog'
);`
	stmt, err := db.Prepare(stmnt)
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

func (i *Ticket) Retrieve(db *sql.DB, id int) (ticket Ticket, err error) {
	ticket = Ticket{}
	err = db.QueryRow("select id, payer_id, due_date from tickets where id = $1", id).Scan(&ticket.Id, &ticket.PayerId, &ticket.DueDate)
	return
}

func (ticket *Ticket) Create(db *sql.DB) (err error) {
	statement := `insert into "tickets" (due_date, payer_id) values ($1, $2) returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ticket.DueDate, ticket.PayerId).Scan(&ticket.Id)
	return err
}

func (ticket *Ticket) Index(db *sql.DB, limit int) ([]Ticket, error) {
	tickets := []Ticket{}
	rows, err := db.Query("select id, due_date, payer_id from tickets limit $1",
		limit)
	if err != nil {
		return []Ticket{}, err
	}
	for rows.Next() {
		ticket := Ticket{}
		err = rows.Scan(&ticket.Id, &ticket.DueDate, &ticket.PayerId)
		if err != nil {
			return []Ticket{}, err
		}
		tickets = append(tickets, ticket)
	}
	rows.Close()
	return tickets, nil
}
