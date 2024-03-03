package main

import "fmt"

func main() {
	a, b, c := "Hello", 12, 42.42
	fmt.Printf("%v -> %T\n", a, a)
	fmt.Printf("%v -> %T\n", b, b)
	fmt.Printf("%v -> %T\n", c, c)
}
