package main

import (
	"html/template"
	"log/slog"
	"path/filepath"

	"github.com/amrojjeh/transactions/internals/models"
)

type templateData struct {
	Form         any
	Transactions []models.Transaction
}

func (app *application) newTemplate() (templateData, error) {
	ts, err := app.tmodel.GetTransactions()
	if err != nil {
		return templateData{}, err
	}
	return templateData{
		Transactions: ts,
	}, nil
}

func (app *application) cacheTemplates() error {
	app.templates = make(map[string]*template.Template)
	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return err
	}

	for _, page := range pages {
		tmpl, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return err
		}

		name := filepath.Base(page)
		app.templates[name], err = tmpl.ParseFiles(page)
		if err != nil {
			return err
		}
		app.logger.Info("template cached", slog.String("page", page))
	}

	return nil
}
