package main

import "fmt"

func f1() {
	fmt.Println("this is f1 function")
}

func f2(a int, b int) {
	fmt.Println("sum:", a+b)
}

func f3(a, b, c int, d, e float64, s string) {
	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "e:", e, "s:", s)
}

func f4(a float64) float64 {
	return a * 2
}

func f5(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	f1()
	f2(5, 7)
	f3(1, 2, 3, 4.5, 6.7, "hello")
	p := f4(1.5)
	fmt.Println("p:", p)
	q, r := f5(10, 5)
	fmt.Println("q:", q, "r:", r)
}
