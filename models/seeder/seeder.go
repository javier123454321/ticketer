package main

import (
	"database/sql"
	"fmt"
	"ticketer/dbconfig"
	"ticketer/models"
)

func main() {
	db := dbconfig.Init()
	seedUsers(db)
	seedTickets(db)
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

func seedTickets(db *sql.DB) {
	tickets := []models.Ticket{
		{PayerId: 1, DueDate: "2023-01-01 10:00:00", Description: "Build a sample application"},
		{PayerId: 1, DueDate: "2023-02-01 12:00:00", Description: "Solve world hunger"},
		{PayerId: 2, DueDate: "2022-12-01 12:00:00", Description: "Get a user flow"},
		{PayerId: 2, DueDate: "2022-12-21 12:00:00", Description: "Web3 sign in"},
	}
	fmt.Printf("seeding %v tickets\n\n", len(tickets))
	for _, ticket := range tickets {
		err := ticket.Create(db)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}
