package routes

import (
	handler "learn_go/test_api/internal/message"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware" // Control of requisitions
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes) // Remove extra '/' at the end of requisitions

	r.Route("/api/user", func(router chi.Router) {
		router.Get("/", handler.HelloWorld)
	})

	// Message struct
	r.Route("/api/json", func(router chi.Router) {
		router.Post("/", handler.ReceiveJSON)
		router.Get("/qtd", handler.QtdMessages)
		router.Get("/get-messages", handler.GetMessages)
		router.Get("/get-message/{code}", handler.GetMessage)
	})

}
