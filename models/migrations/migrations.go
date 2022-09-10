package main

import (
	"ticketer/dbconfig"
	"ticketer/models"
)

func main() {
	db := dbconfig.Init()
	user := models.User{}
	ticket := models.Ticket{}
	user.CreateTable(db)
	ticket.CreateTable(db)
}
