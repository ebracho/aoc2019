package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func fuel(mass int) int {
	if mass <= 6 {
		return 0
	}
	f := (mass / 3) - 2
	return f + fuel(f)
}

func main() {
	in, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf(err.Error())
	}
	result := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf(err.Error())
		}
		result += fuel(mass)
	}
	if scanner.Err() != nil {
		log.Fatalf(scanner.Err().Error())
	}
	fmt.Println(result)
}
