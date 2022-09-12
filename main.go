package main

import (
	"log"
	"net/http"
	"ticketer/controllers"

	_ "github.com/lib/pq"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.ServeFiles("/dist/*filepath", http.Dir("./dist"))
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.Redirect(w, r, "/ticket", http.StatusSeeOther)
	})
	router.GET("/ticket/create", controllers.CreateTicket)
	router.POST("/ticket/create", controllers.StoreTicket)
	router.GET("/ticket/show/:id", controllers.ShowTicket)
	router.GET("/ticket", controllers.IndexTicket)

	log.Fatal(http.ListenAndServe(":8080", router))
}
