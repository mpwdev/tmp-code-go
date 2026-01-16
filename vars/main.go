package main

import "fmt"

func main() {
	name := "Bob"
	fmt.Println("Hello,", name)

	a, b := 10, 20
	fmt.Println("Sum is:", a+b)

	fmt.Printf("Your age is %d\n", 25) // \n add new line symbol

	fmt.Printf("x is %d, y is %d\n", 11, 15)
	fmt.Printf("x is %d, y is %f\n", 20, 30.3) // no new line

	fmt.Printf("He says: \"Hello\"\n")

	radius := 5
	pi := 3.14159

	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("Radius is %-d\n", radius) // note - no minus

	fmt.Printf("Pi is %f\n", pi)
	fmt.Printf("Pi is %+f\n", pi)

	figure := "Circle"

	fmt.Printf("the diameter of a %s with radius %d is %f\n", figure, radius, float64(radius)*2*pi)

	fmt.Printf("This is a %q\n", figure)

	// %v -> replaced by any type
	fmt.Printf("the diameter of a %v with radius %v is %v\n", figure, radius, float64(radius)*2*pi)

	// %T
	fmt.Printf("figure type is %T\n", figure)
	fmt.Printf("radius type is %T\n", radius)
	fmt.Printf("pi type is %T\n", pi)

	// %t
	closed := true
	fmt.Printf("is file closed: %t\n", closed)

	// %b -> base 2
	fmt.Printf("%b\n", 255)
	fmt.Printf("%b\n", 128)
	fmt.Printf("%08b\n", 55)

	// %x
	fmt.Printf("%x\n", 100)

	x := 10.1
	y := 5.7

	fmt.Printf("x * y = %f\n", x*y)
	fmt.Printf("x * y = %0.3f\n", x*y)
}
