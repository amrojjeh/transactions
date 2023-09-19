package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<html>
			<head>
				<title>Hello!</title>
			</head>
			<body>
				<p>Hello World!</p>
			</body>
		</html>
	`))
}

func main() {
	r := mux.NewRouter()
	server := http.Server{
		Handler:      r,
		Addr:         "localhost:8080",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	r.HandleFunc("/", index)

	log.Print("Listening on localhost:8080...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
