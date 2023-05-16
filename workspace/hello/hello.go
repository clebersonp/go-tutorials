package main

import (
	"fmt"
	"golang.org/x/example/stringutil"
)

// run the hello program:
//$ go run example.com/hello

func main() {
	fmt.Println(stringutil.Reverse("Hello"))
	fmt.Println(stringutil.ToUpper("hello"))
}
