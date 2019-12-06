package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func execute(program []int, noun, verb int) int {
	program[1] = noun
	program[2] = verb

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

func search(program []int, target int) int {
	for noun := 0; noun < 100; noun += 1 {
		for verb := 0; verb < 100; verb += 1 {
			tmp := make([]int, len(program))
			copy(tmp, program)
			if execute(tmp, noun, verb) == target {
				return 100*noun + verb
			}

		}
	}
	return -1
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
	fmt.Println(search(program, 19690720))
}
