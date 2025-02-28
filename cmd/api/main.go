package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nikagar4epm/go_api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	var url = "localhost:8000"
	fmt.Println("Listening to requests on", url)
	err := http.ListenAndServe(url, r)

	if err != nil {
		log.Error(err)
	}
}
