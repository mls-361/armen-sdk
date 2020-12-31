/*
------------------------------------------------------------------------------------------------------------------------
####### message ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package message

import (
	"time"

	"github.com/mls-361/uuid"
)

type (
	// Message AFAIRE.
	Message struct {
		Timestamp time.Time
		ID        string
		Topic     string
		Host      string
		Publisher string
		Data      interface{}
	}
)

// New AFAIRE.
func New(topic string, data interface{}) *Message {
	return &Message{
		Timestamp: time.Now(),
		ID:        uuid.New(),
		Topic:     topic,
		Host:      "?",
		Publisher: "?",
		Data:      data,
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
