package main

import (
	"fmt"
	"learn_go/crud_api/internal/handler"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)

	var r *chi.Mux = chi.NewRouter()
	handler.Handler(r)

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
