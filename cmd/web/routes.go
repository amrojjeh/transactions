package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()
	r.Handle("/static/",
		http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))
	r.Handle("/transaction/new", app.new())
	r.Handle("/", app.indexGet())
	return r
}
