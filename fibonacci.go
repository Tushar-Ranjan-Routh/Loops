package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter how many term do you want:")
	fmt.Scanln(&n)
	a := 0
	b := 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		c := a + b
		a = b
		b = c
	}

}
