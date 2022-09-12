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
	router.GET("/ticket/create", createTicket)
	router.POST("/ticket/create", storeTicket)
	router.GET("/ticket/show/:id", ticket)
	router.GET("/ticket", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	files := []string{
		"resources/views/layout.gohtml",
		"resources/views/index.gohtml",
	}
	templates := template.Must(template.ParseFiles(files...))
	db := dbconfig.Init()
	i := models.Ticket{}
	var filter string
	var key string
	if len(r.Form["sortBy"]) != 0 {
		if len(r.Form["key"]) == 0 || r.Form["key"][0] == "all" {
			http.Redirect(w, r, "/ticket", http.StatusPermanentRedirect)
			return
		}
		filter = r.Form["sortBy"][0]
		key = r.Form["key"][0]
	}
	tickets, err := i.Index(db, 15, filter, key)
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
		"resources/views/layout.gohtml",
		"resources/views/ticket.gohtml",
	}
	templates := template.Must(template.ParseFiles(files...))
	ticket := models.Ticket{}
	err = ticket.Retrieve(db, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", ticket)
}

func createTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files := []string{
		"resources/views/layout.gohtml",
		"resources/views/createTicket.gohtml",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", "")
}

func storeTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	form := r.Form
	fmt.Println(form)
	db := dbconfig.Init()
	t := models.Ticket{
		DueDate:     r.Form["due"][0],
		UserId:      1,
		Title:       r.Form["title"][0],
		Description: &r.Form["description"][0],
	}
	if t.Create(db) != nil {
		createTicket(w, r, p)
	}
	route := fmt.Sprintf("/ticket/show/%v", t.Id)
	http.Redirect(w, r, route, http.StatusMovedPermanently)
	fmt.Println("redirected")
}
