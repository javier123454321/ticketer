package models

import (
	"database/sql"
	"fmt"
)

type LineItem struct {
	Id          uint   `json:"id"`
	InvoiceId   uint   `json:"invoice_id"`
	Amount      int64  `json:"amount"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (l *LineItem) CreateTable(db *sql.DB) {
	stmnt := `CREATE TABLE "line_item" (
		id              serial primary key,
		amount      	bigint,
		name	        varchar(255),
		description		text,
		invoice_id      integer references "invoice"(id)
);`
	stmt, err := db.Prepare(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer stmt.Close()
	fmt.Println("creating line_items table")
	_, err = db.Exec(stmnt)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("success!")
}

func (item *LineItem) Create(db *sql.DB) (err error) {
	statement := `insert into "line_item" (amount, name, invoice_id, description) values ($1, $2, $3, $4) returning id`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(item.Amount, item.Name, item.InvoiceId, item.Description).Scan(&item.Id)
	return err
}

func (item *LineItem) GetFromInvoice(db *sql.DB, invoice int) ([]LineItem, error) {
	items := []LineItem{}
	rows, err := db.Query("select id, amount, name, description from line_item where invoice_id = $1",
		invoice)
	if err != nil {
		return []LineItem{}, err
	}
	for rows.Next() {
		item := LineItem{}
		err = rows.Scan(&item.Id, &item.Amount, &item.Name, &item.Description)
		if err != nil {
			return []LineItem{}, err
		}
		items = append(items, item)
	}
	rows.Close()
	return items, nil
}
