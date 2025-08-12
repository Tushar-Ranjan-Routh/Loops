package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Enter two number:")
	fmt.Scan(&a, &b)
	x := a
	y := b
	for b != 0 {
		a, b = b, a%b
	}
	gcd := a
	lcm := x * y / gcd
	fmt.Printf("The gcd of %d and %d is %d.\n", x, y, gcd)
	fmt.Printf("The lcm of %d and %d is %d.\n", x, y, lcm)
}
