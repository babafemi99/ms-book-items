package application

import (
	"github.com/gorilla/mux"
	"log"
	"msitems/client/es"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	es.NewEsclient()
	mapUrls()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Application starting")
	log.Fatal(srv.ListenAndServe())
}
