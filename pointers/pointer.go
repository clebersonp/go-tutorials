package main

import "fmt"

// To declare pointers in golang, we can use the * before the type of data type we want to refer to
// As we can see, the pointer that references an integer is initialized to <nil>.
// We have used * before the data type, this can be anything like *string, *bool, *float64, etc.
// Using the & or the ampersand operator we can get the memory address of a variable
// We have seen that the * is used to declare a pointer variable, but it is also used for de-referencing a pointer.
// So, if we used & to get the memory address of a variable, similarly we can use the * to get back the value from the memory address.
// Both are opposite in terms of accessing the value.

func main() {
	var ptr *int
	fmt.Printf("\nptr value: %v,\nptr type: %T\n\n", ptr, ptr)

	n := 34
	var nPointer *int = &n
	fmt.Printf("n value: %v,\nn type: %T,\nn memory address: %v,\nnPointer value: %v,\nnPointer type: %T,\nnPointer its own memory address %v,\nnPointer de-reference value: %v\n\n", n, n, &n, nPointer, nPointer, &nPointer, *nPointer)
	fmt.Printf("n address memory '%v' is equals to nPointer address memory '%v'? %t\n\n", &n, nPointer, &n == nPointer)
	fmt.Printf("n memory value '%v' is equals to nPointer memory value '%v'? %t\n\n", n, *nPointer, n == *nPointer)

	n = 10
	fmt.Printf("After change the n memory value. n value: %v, *nPointer value: %v\n\n", n, *nPointer)

	*nPointer = 25
	fmt.Printf("After change the *nPointer memory value. n value: %v, *nPointer value: %v\n\n", n, *nPointer)

	// ===========================================================================================================
	fmt.Println("=====================================================================")
	x := 3
	y := 7
	k := &x
	p := &y

	fmt.Printf("Before swapping without pointer: x = %d and y = %d\n\n", x, y)
	a, b := weakSwap(x, y)
	fmt.Printf("After swapping without pointer: x = %d and y = %d\n\n", x, y)
	fmt.Printf("But swapping without pointer return the values swapped: x,a = %d and y,b = %d\n\n", a, b)

	fmt.Printf("Before swapping with pointer: x = %d and y = %d\n\n", x, y)
	swap(k, p)
	fmt.Printf("After swapping with pointer: x = %d and y = %d\n\n", x, y)

	fmt.Println("=====================================================================")

	harryPotter := Book{120, "fiction", "Harry Potter"}
	fmt.Printf("Book value: %v\n\n", harryPotter)
	fmt.Printf("Book address memory: %v\n\n", &harryPotter)
	fmt.Printf("Book.pages address memory: %v\n\n", &harryPotter.pages)
	fmt.Printf("Book.genre address memory: %v\n\n", &harryPotter.genre)
	fmt.Printf("Book.title address memory: %v\n\n", &harryPotter.title)

	fmt.Printf("Book type: %T\n\n", harryPotter)
	fmt.Printf("Book address memory type: %T\n\n", &harryPotter)

	gamesThrones := &harryPotter
	gamesThrones.title = "Games of Thrones"
	fmt.Printf("Harry Potter address: %v, Games Thrones address: %v\n\n", &harryPotter, gamesThrones)

	fmt.Println("=====================================================================")
	book1 := Book{
		pages: 10,
		genre: "Action",
		title: "Robot",
	}
	fmt.Printf("Book value: %v\n\n", book1)
	fmt.Printf("Book address memory: %v\n\n", &book1)
	fmt.Printf("Book.pages address memory: %v\n\n", &book1.pages)
	fmt.Printf("Book.genre address memory: %v\n\n", &book1.genre)
	fmt.Printf("Book.title address memory: %v\n\n", &book1.title)

	book2 := book1
	fmt.Printf("Before change without pointer: Book1: %v, Book2: %v\n\n", book1, book2)
	book2.pages = 25
	book2.title = "Sky"
	book2.genre = "Fiction"
	fmt.Printf("After change without pointer: Book1: %v, Book2: %v\n\n", book1, book2)

}

// swap can exchange the value between the arguments
// x will be with the y value, and vice-versa
func swap(x *int, y *int) {
	*y, *x = *x, *y
}

// weakSwap can exchange the value between the arguments
// x will be with the y value, and vice-versa
// but the exchange occur only with scope local
func weakSwap(x int, y int) (int, int) {
	y, x = x, y
	return x, y
}

// Book represents a book that have properties like pages, genre, title and so on.
type Book struct {
	pages int    // number of pages
	genre string // book's genre
	title string // book's title
}
