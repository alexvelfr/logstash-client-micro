package logstashclientmicro

import "context"

type Client interface {
	LogError(context.Context, Message) error
}

type LogType string

const (
	Error   LogType = "Error"
	Info    LogType = "Info"
	Debug   LogType = "Debug"
	Warning LogType = "Warning"
)

type Message struct {
	Type   LogType `json:"type"`
	XReqID string  `json:"x_req_id"`
	Data   string  `json:"data"`
	File   string  `json:"file"`
	Action string  `json:"action"`
	Error  error   `json:"error"`
}

type message struct {
	Message
	Microservice string `json:"microservice"`
	Error        string `json:"error"`
}
