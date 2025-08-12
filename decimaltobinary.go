package main

import "fmt"

func decimalToBinary(n int) string {
	binary := ""
	for n > 0 {
		r := n % 2
		binary = fmt.Sprintf("%d%s", r, binary)
		n /= 2
	}
	return binary
}

func main() {
	var num int
	fmt.Print("Enter a decimal number: ")
	fmt.Scanln(&num)
	binary := decimalToBinary(num)
	fmt.Printf("Binary of %d is: %s\n", num, binary)
}
