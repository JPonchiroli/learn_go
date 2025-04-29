package main

import (
	"learn_go/templates/pages"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pages.HomePage().Render(r.Context(), w)
	})

	http.ListenAndServe(":8080", nil)
}
