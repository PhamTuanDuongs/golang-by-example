package main

import "fmt"

func defer4() {
	fmt.Println("defer4")
}

func defer5() {
	fmt.Println("defer5")
}

func defer6() {
	fmt.Println("defer6")
}

func defer1() {
	fmt.Println("defer1")
}

func defer2() {
	fmt.Println("defer2")
}

func defer3() {
	fmt.Println("defer3")
}

func checkDefer() {
	defer defer1()
	defer defer2()
	defer defer3()
	defer defer4()
	defer defer5()
	defer defer6()
}
func main() {
	checkDefer()
}
