package main

import (
	"fmt"

	"example.com/tmp/utils"
)

func main() {
	//fmt.Println("Hello, World!")
	//testFunc()
	utils.UtilTestFunc()

	// const
	const c1 = 10
	const c2 = 20.5
	const c3 = "Hello, World!"
	const c4 = true

	// vars types
	var v int = 10
	//var f float64 = 10.5
	f := 10.5
	var s string = "Hello, World!"
	var b bool = true
	// var arr [3]int = [3]int{1, 2, 3}
	// var slice []int = []int{1, 2, 3}
	// var map map[string]int = map[string]int{"apple": 1, "banana": 2}
	// var struct struct {
	// 	name string
	// 	age int
	// } = struct {
	// 	name string
	// 	age int
	// }{"John", 30}
	// var interface interface{} = "Hello, World!"
	// var pointer *int = &v
	// var function func() = func() {
	// 	fmt.Println("Hello, World!")
	// }

	fmt.Println("v int:", v)
	fmt.Println("f float64:", f)
	fmt.Println("s string:", s)
	fmt.Println("b bool:", b)
}
