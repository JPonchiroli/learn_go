package message

import (
	"encoding/json"
	"fmt"
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

	MessagesMap[message.Code] = message

	w.Write([]byte("Message Saved"))
}

func QtdMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Qtd Messages: " + strconv.Itoa(qtdMessages)))
}

func GetMessages(w http.ResponseWriter, r *http.Request) {
	for code, message := range MessagesMap {
		response := fmt.Sprintf("The code is %v and the message is %v\n", code, message.Message)
		w.Write([]byte(response))
	}
}
