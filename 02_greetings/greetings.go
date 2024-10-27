package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message
	message := fmt.Sprintf("Hi, %v, Welcome!", name);
	return message, nil
}

//randomFormat returns one of a set of greeting messages
//The returned message is selected at random
