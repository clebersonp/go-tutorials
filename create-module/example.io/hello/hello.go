// Declare a main package. In Go, code executed as an application must be in a main package.
package main

import (
	"create-module/example.io/greetings"
	"fmt"
	"log"
)

// The go build command compiles the packages, along with their dependencies, but it doesn't install the results.
// The go install command compiles and installs the packages.
// $ go build command to compile the code into an executable.
// $ go install, build and install the package (move) the executable file to the user directory and add to the PATH, so just type hello from anywhere in the terminal to run the program: Ex: my_user_path\go\bin\hello.exe
// $ go env, to list all go env information like GOPATH=my_user_path\go
func main() {
	// Set properties of the predefined Logger, including the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin", "John", "Maria", "Carlos", "Allan", "Anna", "Melissa"}

	// Request a greeting message.
	//message, err := greetings.Hello("Gladys")

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)

	// If an error was returned, print it io console and exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned message to the console.
	//fmt.Println(message)

	// If no error was returned, print the returned map of messages to the console.
	fmt.Println(messages)
}
