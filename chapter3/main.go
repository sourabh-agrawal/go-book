package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1.0+1.0)

	// String type
	fmt.Println(`Hello 
	world\n`)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")

	// Boolean type
	fmt.Println()
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(false && true)
	fmt.Println()
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(false || true)
	fmt.Println()
	fmt.Println(!true)
	fmt.Println(!false)
}
