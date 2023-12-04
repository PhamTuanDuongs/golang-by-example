package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := &Person{Name: "John", Age: 30}
	fmt.Println("Original:", *person)

	// Set the pointer to the struct to nil
	person = nil
	fmt.Println("After deletion:", person)
}
