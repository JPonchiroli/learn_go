package message

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var MessagesMap = make(map[int]Message)
