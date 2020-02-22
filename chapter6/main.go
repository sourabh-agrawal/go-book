package main

import "fmt"

func main() {
	fmt.Println("Chapter 6")

	// Array
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)

	// Improved version of above
	a := [5]float64{98, 93, 77, 82, 83}
	total = 0
	for _, value := range a {
		total += value
	}
	fmt.Println(total / float64(len(a)))

	// Let the compiler determine the size
	arr := [...]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// Slice

	// y := []int{1, 2, 3, 4, 5}
	s := make([]float64, 5, 10)
	for i := 0; i < 5; i++ {
		s[i] = 1
	}
	fmt.Println(s)

	s = a[:]
	fmt.Println(s)

	s = a[:3]
	fmt.Println(s)

	s = a[1:]
	fmt.Println(s)

	s = a[1:3]
	fmt.Println(s)

	// Slice functions
	// append
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	// copy
	slice3 := make([]int, 2)
	copy(slice3, slice2)
	fmt.Println(slice2, slice3)

	// map
	m := make(map[string]string)
	m["name"] = "sourabh"
	fmt.Println(m)

	m = map[string]string{
		"name": "sourabh",
	}
	fmt.Println(m)

	keyboards := map[string]map[string]string{
		"key1": map[string]string{
			"brand":  "Keychron",
			"model":  "k6",
			"frame":  "Aluminium",
			"switch": "Gateron Optical Brown",
			"type":   "swappable",
		},
		"key2": map[string]string{
			"brand":  "Geek",
			"model":  "GK61",
			"frame":  "plastic",
			"switch": "Gateron Optical Brown",
			"type":   "swappable",
		},
	}
	fmt.Println(keyboards["key1"])

	if keyboard, ok := keyboards["key2"]; ok {
		fmt.Println(keyboard)
	}
}
