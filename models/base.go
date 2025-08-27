package models

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	MessageDetail string `json:"message_detail"`
}
