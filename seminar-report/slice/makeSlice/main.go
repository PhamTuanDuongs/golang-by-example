package main

import "fmt"

func main() {
	// Slices are declared with syntax similar to array. Only length is not specified
	cities := []string{"Ha Noi", "Ninh Binh", "Nam Dinh"}

	// declare a slice from an array
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]
	slice = arr[:5]
	slice = arr[1:5]

	// add elements to a slice
	slice = append(slice, 6)
	fmt.Println(cities, slice)
	sliceFromFunc()
}

func sliceFromFunc() {
	slice := make([]int, 3, 6)
	fmt.Println(&slice[0])
	fmt.Println(slice)
	for i := 0; i <= 1024; i++ {
		slice = append(slice, i)
	}
	fmt.Println(cap(slice))
}
