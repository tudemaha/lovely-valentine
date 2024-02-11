package model

import (
	"log"
)

func CheckAvailableID(id string) (Message, bool) {
	var message Message

	message, err := GetMessage(id)
	if err != nil {
		log.Printf("ERROR CreateMessage fatal error: %v", err)
		return message, false
	}

	return message, true
}
