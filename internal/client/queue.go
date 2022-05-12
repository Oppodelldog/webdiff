package client

import (
	"encoding/json"
	"fmt"
)

type QueueUpdate struct {
	Queue int `json:"queue"`
}

func NotifyQueueUpdate(queueSize int) {
	if hub == nil {
		fmt.Printf("client.NotifyQueueUpdate was called, but hub is not initialized yet\n")

		return
	}

	jsonMessage, err := json.Marshal(QueueUpdate{Queue: queueSize})
	if err != nil {
		fmt.Printf("client.NotifyQueueUpdate could not marshal message: %v\n", err)

		return
	}

	hub.Broadcast(jsonMessage)
}
