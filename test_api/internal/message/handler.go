package message

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var qtdMessages int = 0

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func ReceiveJSON(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)

	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	qtdMessages++

	w.Write([]byte("Receive JSON: " + message.Message))
}

func QtdMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Qtd Messages: " + strconv.Itoa(qtdMessages)))
}
