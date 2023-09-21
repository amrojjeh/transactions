package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

type application struct {
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP Network Address")
	flag.Parse()
	app := application{}
	server := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("listening on %s", *addr)

	err := server.ListenAndServe()
	log.Fatal(err)
}
