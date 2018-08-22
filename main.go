package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var x float64 = 20.0

	var a, b, c = 3, 4, "foo"

	y := 42

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
