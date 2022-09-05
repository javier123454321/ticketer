package main

import (
	"sample-invoicer/dbconfig"
	"sample-invoicer/models"
)

func main() {
	db := dbconfig.Init()
	user := models.User{}
	invoice := models.Invoice{}
	items := models.LineItem{}
	user.CreateTable(db)
	invoice.CreateTable(db)
	items.CreateTable(db)
}
