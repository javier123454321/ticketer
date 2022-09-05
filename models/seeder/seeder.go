package main

import (
	"database/sql"
	"fmt"
	"sample-invoicer/dbconfig"
	"sample-invoicer/models"
)

func main() {
	db := dbconfig.Init()
	seedUsers(db)
	seedInvoices(db)
	seedLineItems(db)
}

func seedUsers(db *sql.DB) {
	u := []models.User{
		{Name: "Bob James", Email: "bob@bobjames.com"},
		{Name: "James Roberts", Email: "james@jamesroberts.com"},
		{Name: "Caligula Roman", Email: "cali@gmail.com"},
	}
	fmt.Printf("seeding %v users\n\n", len(u))
	for _, user := range u {
		err := user.Create(db)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

func seedInvoices(db *sql.DB) {
	invoices := []models.Invoice{
		{PayerId: 1, DueDate: "2023-01-01 10:00:00"},
		{PayerId: 1, DueDate: "2023-02-01 12:00:00"},
		{PayerId: 2, DueDate: "2022-12-01 12:00:00"},
		{PayerId: 2, DueDate: "2022-12-21 12:00:00"},
		{PayerId: 2, DueDate: "2022-12-10 12:00:00"},
	}
	fmt.Printf("seeding %v invoices\n\n", len(invoices))
	for _, invoice := range invoices {
		err := invoice.Create(db)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

func seedLineItems(db *sql.DB) {
	items := []models.LineItem{
		{InvoiceId: 1, Amount: 2500000000000, Name: "Services Rendered", Description: "For building the sample application"},
		{InvoiceId: 1, Amount: 400000000000, Name: "Expenditure Reunbursement", Description: "Restaurants, etc"},
		{InvoiceId: 1, Amount: 15000000000000, Name: "Lorem ipsum dolor", Description: "ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat"},
		{InvoiceId: 2, Amount: 300000000000, Name: "quasi architecto", Description: "ugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, "},
		{InvoiceId: 2, Amount: 25000000000000, Name: "tempor incididunt", Description: "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the syste"},
		{InvoiceId: 2, Amount: 2000000000000, Name: "nulla pariatur", Description: "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias"},
		{InvoiceId: 3, Amount: 2500000000000, Name: "Services Rendered", Description: "For building the sample application"},
		{InvoiceId: 3, Amount: 400000000000, Name: "Expenditure Reunbursement", Description: "Restaurants, etc"},
		{InvoiceId: 4, Amount: 15000000000000, Name: "Lorem ipsum dolor", Description: "ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat"},
		{InvoiceId: 5, Amount: 300000000000, Name: "quasi architecto", Description: "ugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, "},
		{InvoiceId: 2, Amount: 25000000000000, Name: "tempor incididunt", Description: "But I must explain to you how all this mistaken idea of denouncing pleasure and praising pain was born and I will give you a complete account of the syste"},
		{InvoiceId: 1, Amount: 2000000000000, Name: "nulla pariatur", Description: "At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti quos dolores et quas molestias"},
	}
	fmt.Printf("seeding %v line items\n\n", len(items))
	for _, item := range items {
		err := item.Create(db)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
