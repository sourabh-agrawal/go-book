package main

import "fmt"

var (
	no     float64
	output float64
)

func main() {
	var x string
	x = "Hello"
	fmt.Println(x)
	x += " World"
	fmt.Println(x)

	// String comparison
	var a string = "Hello"
	var b string = "Hello"
	fmt.Println(a == b)

	// we can also initialize two variables together
	var h, j int = 3, 4
	fmt.Println(h, j)

	// shorhands
	z := "Hello World"
	fmt.Println(z)

	var y = "Hello World"
	fmt.Println(y)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	// constants
	const c = "Hello World"
	fmt.Println(c)

	// Example program
	fmt.Print("Enter a number: ")
	fmt.Scanf("%f", &no)
	output = no * 2
	fmt.Println("You entered: ", no, " and the double of it is ", output)

}
