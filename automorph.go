package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number:")
	fmt.Scan(&n)
	sq := n * n
	temp := n
	m := 1
	for n > 0 {
		m *= 10
		n /= 10
	}
	if sq%m == temp {
		fmt.Println("It is an automorph number.")
	} else {
		fmt.Println("It is not an automorph number.")
	}
}
