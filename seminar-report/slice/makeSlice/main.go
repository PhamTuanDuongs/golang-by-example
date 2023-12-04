package main

import "fmt"

func main() {
	slice := make([]int, 3, 6)
	fmt.Println(&slice[0])
	fmt.Println(slice)
	for i := 0; i <= 1024; i++ {
		slice = append(slice, i)
	}
	fmt.Println(cap(slice))
}
