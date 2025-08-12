package main

import "fmt"

func digitcount(a int) int {
	c := 0
	for a > 0 {
		c++
		a /= 10
	}
	return c
}
func product(a, b int) int {
	r := 1
	for i := 1; i <= b; i++ {
		r *= a
	}
	return r
}
func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scan(&n)
	temp := n
	s := 0
	d := digitcount(n)
	for n > 0 {
		r := n % 10
		s += product(r, d)
		n /= 10
	}
	if temp == s {
		fmt.Println("It is an armstrong number.")
	} else {
		fmt.Println("It is not an armstrong number.")
	}
}
