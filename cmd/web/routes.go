package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", app.index)
	return r
}

func (app *application) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/pages/transactions.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		log.Printf(err.Error())
		return
	}

	buff := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buff, "base", nil)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		log.Printf(err.Error())
		return
	}

	buff.WriteTo(w)
}
