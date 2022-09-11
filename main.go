package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"ticketer/dbconfig"
	"ticketer/models"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/dist/*filepath", http.Dir("./dist"))
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.Redirect(w, r, "/ticket", 301)
	})
	router.GET("/ticket/:id", ticket)
	router.GET("/ticket", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files := []string{
		"resources/views/layout.html",
		"resources/views/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	db := dbconfig.Init()
	i := models.Ticket{}
	tickets, err := i.Index(db, 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", tickets)
}

func ticket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
	}
	db := dbconfig.Init()
	files := []string{
		"resources/views/layout.html",
		"resources/views/ticket.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	ticket := models.Ticket{}
	err = ticket.Retrieve(db, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", ticket)
}
