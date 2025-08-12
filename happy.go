package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number:")
	fmt.Scanln(&n)
	var s int
	for {
		s = 0
		for n > 0 {
			r := n % 10
			s += r * r
			n /= 10
		}
		if s >= 10 {
			n = s
			continue
		} else {
			break
		}
	}
	if s == 1 {
		fmt.Println("It is a happy number.")
	} else {
		fmt.Println("It is no a happy number.")
	}
}
