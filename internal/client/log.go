package client

import (
	"encoding/json"
	"fmt"
	"github.com/Oppodelldog/webdiff/pkg/ws"
)

type LogMessage struct {
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

var hub *ws.Hub

func StartWebsocketHub() *ws.Hub {
	hub = ws.StartHub()

	return hub
}

func Log(message, severity string) {
	if hub == nil {
		fmt.Printf("client.Log was called, but hub is not initialized yet: message='%s'\n", message)

		return
	}

	jsonMessage, err := json.Marshal(LogMessage{Message: message, Severity: severity})
	if err != nil {
		fmt.Printf("client.Log could not marshal message: %v\n", err)

		return
	}

	hub.Broadcast(jsonMessage)
}
