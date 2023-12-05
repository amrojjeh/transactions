package main

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.logger.Error("server error", slog.String("error", err.Error()))
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func (app *application) render(w http.ResponseWriter, page string,
	data templateData, status int) {
	tmpl, exists := app.templates[page]

	if !exists {
		app.serverError(w, errors.New(fmt.Sprintf("page %v does not exist", page)))
		return
	}

	buff := new(bytes.Buffer)
	err := tmpl.ExecuteTemplate(buff, "base", data)

	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	buff.WriteTo(w)
}
