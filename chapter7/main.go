package main

import "fmt"

func main() {
	fmt.Println("Chapter7: Functions")

	numbers := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(numbers))

	// Example of named return type
	fmt.Println(lengthOfName("sourabh"))

	// Example of multiple return values
	name, length := nameProperties("sourabh")
	fmt.Println("Length of ", name, " is ", length)

	// Example of multiple return values with named return type
	name, length = namePropertiesWithNamedReturnType("sourabh")
	fmt.Println("Length of ", name, " is ", length)

	// Example of variadic function
	fmt.Println("Average of numbers from 1 to 5 is ", avg(1, 2, 3, 4, 5))
	// We can also use ... while passing slice to function
	fmt.Println("Average of numbers is ", avg(numbers...))

	// Closure
	// Print even numbers for each call
	printEven := printEvenNumbers()
	fmt.Println(printEven()) // 0
	fmt.Println(printEven()) // 2

	// Recursion
	// Calculate factorial
	fmt.Print("Enter no: ")
	var no uint
	fmt.Scanf("%d", &no)
	fmt.Println("factorial of entered number is ", factorial(no))

	// defer
	// defer is called after function completes
	defer goodBye()

	// panic and recover
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("exit")
}

func average(s []float64) float64 {
	total := 0.0
	for _, value := range s {
		total += value
	}

	return total / float64(len(s))
}

func lengthOfName(name string) (length int) {
	length = len(name)
	return
}

func nameProperties(name string) (string, int) {
	return name, len(name)
}

func namePropertiesWithNamedReturnType(n string) (name string, length int) {
	name = n
	length = len(name)
	return
}

func avg(no ...float64) float64 {
	total := 0.0
	for _, value := range no {
		total += value
	}
	return total / float64(len(no))
}

func printEvenNumbers() func() int {
	i := 0
	return func() (no int) {
		no = i
		i += 2
		return
	}
}

func factorial(no uint) uint {
	if no == 0 {
		return 1
	}
	return no * factorial(no-1)
}

func goodBye() {
	fmt.Println("GoodBye")
}
