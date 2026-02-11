package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 goroutine execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i =", i)
	}
	fmt.Println("f1 goroutine execution ended")
}

func f2() {
	fmt.Println("f2 goroutine execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i =", i)
	}
	fmt.Println("f2 goroutine execution ended")
}

func main() {
	fmt.Println("main execution started")
	go f1()
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	f2()
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)
	fmt.Println("No of Goroutines:", runtime.NumGoroutine())
	fmt.Println("main execution ended")

}
