package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scanln(&n)
	temp := n
	s := 0
	for n > 0 {
		r := n % 10
		s += r
		n /= 10
	}
	if temp%s == 0 {
		fmt.Println("It is a niven number.")
	} else {
		fmt.Println("It is not a niven number.")
	}
}
