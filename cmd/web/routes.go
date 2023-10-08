package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	r.Handle("/", app.indexGet())
	r.Handle("/transaction/new", app.newGet())
	return r
}
