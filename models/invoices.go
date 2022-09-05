package models

import (
	"database/sql"
	"fmt"
)

type Invoice struct {
	Id        uint       `json:"id"`
	PayerId   uint       `json:"payer_id"`
	DueDate   string     `json:"due_date"`
	LineItems []LineItem `json:"line_items"`
}

func (i *Invoice) CreateTable(db *sql.DB) {
	stmnt := `CREATE TABLE "invoice" (
		id              serial primary key,
		due_date        timestamp,
		payer_id        integer references "user"(id)
);`
	stmt, err := db.Prepare(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	fmt.Println("creating invoices table")
	_, err = db.Exec(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success!")
}

func (i *Invoice) Retrieve(db *sql.DB, id int) (invoice Invoice, err error) {
	invoice = Invoice{}
	err = db.QueryRow("select id, payer_id, due_date from invoice where id = $1", id).Scan(&invoice.Id, &invoice.PayerId, &invoice.DueDate)
	return
}

func (invoice *Invoice) Create(db *sql.DB) (err error) {
	statement := `insert into "invoice" (due_date, payer_id) values ($1, $2) returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(invoice.DueDate, invoice.PayerId).Scan(&invoice.Id)
	return err
}

func (invoice *Invoice) Index(db *sql.DB, limit int) ([]Invoice, error) {
	invoices := []Invoice{}
	rows, err := db.Query("select id, due_date, payer_id from invoice limit $1",
		limit)
	if err != nil {
		return []Invoice{}, err
	}
	for rows.Next() {
		invoice := Invoice{}
		err = rows.Scan(&invoice.Id, &invoice.DueDate, &invoice.PayerId)
		if err != nil {
			return []Invoice{}, err
		}
		invoices = append(invoices, invoice)
	}
	rows.Close()
	return invoices, nil
}
