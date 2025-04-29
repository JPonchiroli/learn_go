package main

import (
	"fmt"
	"learn_go/test_api/internal/routes"
	"net/http"

	"github.com/go-chi/chi"          // Roting http
	log "github.com/sirupsen/logrus" // Generate logs
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter() // Receive the http requests and direct them to the handler
	routes.Handler(r)

	fmt.Println("Starting CRUD API service...")

	fmt.Println(`
	 ______     ______        ______     ______   __    
	/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
	\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
	 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
	  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ `)

	err := http.ListenAndServe("localhost:8080", r)

	if err != nil {
		log.Error(err)
	}
}
