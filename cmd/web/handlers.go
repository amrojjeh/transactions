package main

import (
	"net/http"
)

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	app.render(w, "transactions.tmpl", templateData{}, http.StatusOK)
}
