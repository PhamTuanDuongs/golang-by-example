package main

import "fmt"

type HigherFunc func() int
type Human interface{}

func main() {

	var _ *int = nil
	var _ Human = nil
	var _ map[int]string = nil
	var _ []int = nil
	var _ chan int = nil
	var _ chan string = nil

	var _ map[string]int = nil
	nil := 123
	fmt.Println(nil) // 123

}
