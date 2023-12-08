package main

import (
	"fmt"
	"sync"
)

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

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ", r)
		}
	}()
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < 5; i++ {
		fmt.Println("f1", arr[i])
	}

}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in ", r)
		}
	}()
	var ptr *int
	fmt.Println(*ptr)
}
func main() {
	var wg sync.WaitGroup
	checkDefer()
	wg.Add(2)
	go f1(&wg)
	go f2(&wg)
	wg.Wait()
	fmt.Println("Can touch this statement")
	result := exampleFunction()
	fmt.Println(result)
}

func exampleFunction() (result int) {
	// The "defer" keyword is used to schedule a function to be executed later, after the surrounding function returns.
	defer func() {
		// Inside the deferred function, we can read and assign values to the named return values of the enclosing function.
		result = result * 2
	}()

	// Assign a value to the named return value
	result = 42

	// The surrounding function will return the value assigned to "result"
	return result
}
