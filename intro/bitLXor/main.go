//Provided 2 numbers a and b, return value of the LXOR operation, opposite to regural bit by bit XOR
package main

import (
	"fmt"
)

func main() {
	fmt.Println(bitLXor(15, 3))
}

func bitLXor(a, b int) int {
	t, u := a, b
	fmt.Printf("a: %08b\n", a)
	fmt.Printf("b: %08b\n", b)
	for t != u {
		fmt.Println("t, u: ", t, u)
		fmt.Printf("t: %08b\n", t)
		fmt.Printf("u: %08b\n", u)
		if t == 0 {
			a *= 2
		}
		if u == 0 {
			b *= 2
		}
		t /= 2
		u /= 2
		fmt.Println("a, b:\n", a, b)
		fmt.Printf("a: %08b\n", a)
		fmt.Printf("b: %08b\n", b)
	}
	fmt.Println("a, b:\n", a, b)
	fmt.Printf("a: %08b\n", a)
	fmt.Printf("b: %08b\n", b)
	return a ^ b
}
