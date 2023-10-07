package main

import (
	"html/template"
	"log/slog"
	"path/filepath"
)

type templateData struct {
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
