package res

type response struct {
	Message     string      `json:"message"`
	MessageType string      `json:"message_type"`
	Data        interface{} `json:"data"`
}

func Err(message string, data interface{}) *response {
	return &response{Message: message, MessageType: "ERROR", Data: data}
}

func Success(message string, data interface{}) *response {
	return &response{Message: message, MessageType: "Success", Data: data}
}
