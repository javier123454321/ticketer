package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sample-invoicer/dbconfig"
	"sample-invoicer/models"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.Redirect(w, r, "/invoice", 301)
	})
	router.GET("/invoice/:id", invoice)
	router.GET("/invoice", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files := []string{
		"views/layout.html",
		"views/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	db := dbconfig.Init()
	i := models.Invoice{}
	invoices, err := i.Index(db, 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", invoices)
}

func invoice(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	db := dbconfig.Init()
	files := []string{
		"views/layout.html",
		"views/invoice.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	invoice := models.Invoice{}
	info, err := invoice.Retrieve(db, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	items := models.LineItem{}
	lineItems, err := items.GetFromInvoice(db, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	info.LineItems = lineItems
	templates.ExecuteTemplate(w, "layout", info)
}
