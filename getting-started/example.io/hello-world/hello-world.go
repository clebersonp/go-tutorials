// Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
package main

// Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
// Add new module requirements and sums. Go will add the quote module as a requirement, as well as a go.sum file for use in authenticating the module. $ go mod tidy
// When you ran go mod tidy, it located and downloaded the rsc.io/quote module that contains the package you imported. By default, it downloaded the latest version -- v1.5.2.
import (
	"fmt"
	"rsc.io/quote"
)

// Implement a main function to print a message to the console. A main function executes by default when you run the main package.
// $ go run .
// you can build and install that program with the go tool: go install example/hello
// This command builds the hello command, producing an executable binary. It then installs that binary as $HOME/go/bin/hello (or, under Windows, %USERPROFILE%\go\bin\hello.exe).
// Declare a main package. In Go, code executed as an application must be in a main package.
func main() {
	fmt.Println("Hello, World!!!")
	fmt.Println(quote.Go())
}
