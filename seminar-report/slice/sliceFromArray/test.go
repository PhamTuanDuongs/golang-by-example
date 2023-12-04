package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4}

	slices := arr[:]

	fmt.Println(cap(slices))

	slices = append(slices, 6)
	slices[0] = 10
	fmt.Println(cap(slices))

	fmt.Println(slices)

	fmt.Println(&slices[0])
	fmt.Println(&arr[0])
	fmt.Println(arr[0])
}
