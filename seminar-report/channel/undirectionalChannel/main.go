package main

import (
	"fmt"
	"sync"
)

// func 1

// func 2

// func main
// defer func1
// defer func2

func sendData(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

// Hàm nhận dữ liệu từ kênh (receive-only channel)
func receiveData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Kiểm tra xem kênh đã đóng chưa
		data, ok := <-ch
		if !ok {
			fmt.Println("Channel closed.")
			return
		}
		fmt.Println("Received:", data)
	}
}

func main() {
	// Tạo một kênh
	myChannel := make(chan int)

	// Sử dụng WaitGroup để đợi tất cả goroutine hoàn thành
	var wg sync.WaitGroup

	// Khởi động goroutine gửi dữ liệu
	wg.Add(1)
	go sendData(myChannel, &wg)

	// Khởi động goroutine nhận dữ liệu
	wg.Add(1)
	go receiveData(myChannel, &wg)

	// Đợi tất cả goroutine hoàn thành
	wg.Wait()
}
