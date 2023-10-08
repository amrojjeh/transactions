package main

import (
	"net/http"
)

func (app *application) indexGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		app.render(w, "transactions.tmpl", templateData{}, http.StatusOK)
	})

}

func (app *application) newGet() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render(w, "new.tmpl", templateData{}, http.StatusOK)
	})
}
