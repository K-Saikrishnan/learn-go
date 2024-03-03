package main

import "fmt"

func main() {
	const (
		a = 747
		b = 911
		c = 90210
	)

	fmt.Printf("%d \t %b \t\t %x\n", a, a, a)
	fmt.Printf("%d \t %b \t\t %x\n", b, b, b)
	fmt.Printf("%d \t %b \t %x\n", c, c, c)
}
