package main

import (
	"fmt"
)

func digit(n, place int) int {
	return (n / place) % 10
}

func isSolution(i int) bool {
	adj := false
	for ; i > 0; i /= 10 {
		if digit(i, 1) == digit(i, 10) {
			adj = true
		}
		if digit(i, 1) < digit(i, 10) {
			return false
		}
	}
	return adj
}

func main() {
	res := 0
	for i := 136818; i <= 685979; i += 1 {
		if isSolution(i) {
			res += 1
		}
	}
	fmt.Println(res)
}
