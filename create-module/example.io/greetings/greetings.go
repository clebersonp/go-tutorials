package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Hello returns a greeting for the named person.
// In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
// Ref: https://go.dev/doc/tutorial/create-module
// Change the function so that it returns two values: a string and an error. Your caller will check the second value to see if an error occurred.
// Any Go function can return multiple values.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if strings.TrimSpace(name) == "" {
		return "", errors.New("empty name")
	}

	/*
		In Go, the := operator is a shortcut for declaring and initializing a variable in one line
		(Go uses the value on the right to determine the variable's type). Taking the long way, you might have written this as:

		var message string
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	// Return a greeting that embeds the name in a message
	// The first argument is a format string, and Sprintf substitutes the name parameter's value for the %v format verb. Inserting the value of the name parameter completes the greeting text.
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)

	// Add nil (meaning no error) as a second value in the successful return. That way, the caller can see that the function succeeded
	// nil is like null in java
	return message, nil
}

// init sets initial values for variables used in the function.
// Add an init function to seed the rand package with the current time.
// Go executes init functions automatically at program startup, after global variables have been initialized
func init() {
	rand.NewSource(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message is selected at random.
// Note that randomFormat starts with a lowercase letter, making it accessible only to code in its own package (in other words, it's not exported).
func randomFormat() string {
	// A slice of message formats.
	// When declaring a slice, you omit its size in the brackets, like this: []string. This tells Go that the size of the array underlying the slice can be dynamically changed
	formats := []string{
		"Hi, %v. Welcome!!",
		"Great to see you, %v!!",
		"Hail, %v! Well met!!",
	}

	// Return a randomly selected message format by specifying a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	// In Go, you initialize a map with the following syntax: make(map[key-type]value-type)
	messages := make(map[string]string)

	// Loop through the received slice of names, calling the Hello function to get a message for each name.
	//  In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value.
	// You don't need the index, so you use the Go blank identifier (an underscore) to ignore it.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with the name.
		messages[name] = message
	}
	return messages, nil
}
