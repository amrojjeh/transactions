package main

import (
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger    *slog.Logger
	templates map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := application{
		logger: logger,
	}
	err := app.cacheTemplates()
	if err != nil {
		logger.Error("couldn't cache templates", slog.String("error", err.Error()))
		os.Exit(1)
	}
	server := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	logger.Info("server is listening", slog.String("addr", *addr))

	err = server.ListenAndServe()
	logger.Error("server error", slog.String("error", err.Error()))
	os.Exit(1)
}
