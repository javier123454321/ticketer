package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"ticketer/dbconfig"
	"ticketer/models"

	"github.com/julienschmidt/httprouter"
)

func IndexTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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
		panic(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", tickets)
}

func ShowTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		panic(err.Error())
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
		panic(err.Error())
	}
	templates.ExecuteTemplate(w, "layout", ticket)
}

func CreateTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	files := []string{
		"resources/views/layout.gohtml",
		"resources/views/createTicket.gohtml",
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", "")
}

func StoreTicket(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	db := dbconfig.Init()
	t := models.Ticket{
		DueDate:     r.Form["due"][0],
		UserId:      1,
		Title:       r.Form["title"][0],
		Description: &r.Form["description"][0],
	}
	if t.Create(db) != nil {
		CreateTicket(w, r, p)
	}
	route := fmt.Sprintf("/ticket/show/%v", t.Id)
	http.Redirect(w, r, route, http.StatusSeeOther)
}
