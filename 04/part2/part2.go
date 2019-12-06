package main

import (
	"fmt"
)

func digit(n, place int) int {
	return (n / place) % 10
}

func isSolution(i int) bool {
	run := 1
	twoRunFound := false
	for ; i > 0; i /= 10 {
		ones := digit(i, 1)
		tens := digit(i, 10)
		if ones == tens {
			run += 1
		} else {
			if run == 2 {
				twoRunFound = true
			}
			run = 1
		}
		if ones < tens {
			return false
		}
	}
	return twoRunFound
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
