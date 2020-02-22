package main

import "fmt"

func main() {
	fmt.Println("Chapter8: pointers")

	// First way to use pointers
	x := 5
	one(&x)
	fmt.Println(x)

	// Second way is to use new
	y := new(int)
	one(y)
	// y is a pointer, so derefrence it to get the value
	fmt.Println(*y)

	z := 1.5
	square(&z)
	fmt.Println(z)

	a := 1
	b := 2
	fmt.Println(a, b)
	swap(&a, &b)
	fmt.Println(a, b)
}

func one(ptr *int) {
	*ptr = 0
}

func square(ptr *float64) {
	*ptr = *ptr * *ptr
}

func swap(ptr1 *int, ptr2 *int) {
	// let's store the value that ptr1 reference to in s temporary variable
	temp := *ptr1
	*ptr1 = *ptr2
	*ptr2 = temp
}
