package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/amrojjeh/transactions/internals/models"
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger    *slog.Logger
	templates map[string]*template.Template
	tmodel    models.TransactionModel
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := sql.Open("mysql", "web:pass@/BEApp?parseTime=true")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("could not ping mysql db", slog.String("error", err.Error()))
		os.Exit(1)
	}
	if err != nil {
		logger.Error("could not open mysql db", slog.String("error", err.Error()))
		os.Exit(1)
	}
	app := application{
		logger: logger,
		tmodel: models.TransactionModel{
			DB: db,
		},
	}
	err = app.cacheTemplates()
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
