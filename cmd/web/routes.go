package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("/", app.index)
	return r
}
