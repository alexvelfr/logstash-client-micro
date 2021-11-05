package logstashclientmicro

import "context"

type Client interface {
	LogError(context.Context, Message) error
}

type Message struct {
	XReqID string `json:"x_req_id"`
	Data   string `json:"data"`
	File   string `json:"file"`
	Action string `json:"action"`
	Error  error  `json:"error"`
}

type message struct {
	Message
	Microservice string `json:"microservice"`
	Error        string `json:"error"`
}
