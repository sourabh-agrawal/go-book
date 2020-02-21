package main

import "fmt"

func main() {
	var i int = 1
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}

	// Second way
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}

	// odd even check
	fmt.Println()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			// even
			fmt.Println(i, "even")
		} else {
			// odd
			fmt.Println(i, "odd")
		}
	}

	// check if number is divisible by 2, 3 or 5
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "divisible by 2")
		} else if i%3 == 0 {
			fmt.Println(i, "divisible by 3")
		} else if i%5 == 0 {
			fmt.Println(i, "divisible by 5")
		} else {
			fmt.Println("default case")
		}
	}

	// switch case
	fmt.Print("Enter a number")
	fmt.Scanf("%d", &i)
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	case 9:
		fmt.Println("Nine")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("Entered number is not in range of 0..10")
	}
}
