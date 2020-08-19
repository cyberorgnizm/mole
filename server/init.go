package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Initialize starts up a http server listerning on given port
// and creates a router for the index route path
func Initialize() {
	const addr = "127.0.0.1"
	port := flag.String("p", "8000", "Port for HTTP server")
	flag.Parse()

	router := mux.NewRouter()
	hub := newHub()
	go hub.run()

	// handles web socket connection
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	applicationHandler := Application{staticPath: "build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(applicationHandler)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%s", addr, *port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("[Server] Running on %s:%s...", addr, *port)
	log.Fatal(srv.ListenAndServe())
}
