package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func execute(program []int) int {
	program[1] = 12
	program[2] = 2

	i := 0
	for {
		switch op := program[i]; op {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case 99:
			return program[0]
		default:
			panic(fmt.Sprintf("invalid opcode %d (program state = %v)", op, program))
		}
		i += 4
	}
}

func main() {
	in, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	elements := strings.Split(strings.TrimSpace(string(in)), ",")
	program := make([]int, 0, len(elements))
	for _, s := range elements {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err.Error())
		}
		program = append(program, i)
	}
	fmt.Println(execute(program))
}
