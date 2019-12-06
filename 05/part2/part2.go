package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type getter func(int) int

func input() int {
	return 5
}

func execute(program []int) int {
	pos := func(i int) int { return program[i] }
	val := func(i int) int { return i }
	getters := map[int]getter{0: pos, 1: val}

	i := 0
	for {

		instr := program[i]

		first := getters[(instr/100)%10](min(len(program)-1, i+1))
		second := getters[(instr/1000)%10](min(len(program)-1, i+2))
		third := getters[(instr/10000)%10](min(len(program)-1, i+3))

		jumped := false

		// fmt.Printf("%d %d %v\n", i, instr, program)

		switch op := instr % 100; op {
		case 1:
			program[third] = program[first] + program[second]
		case 2:
			program[third] = program[first] * program[second]
		case 3:
			program[first] = input()
		case 4:
			fmt.Println(program[first])
		case 5:
			if program[first] != 0 {
				jumped = true
				i = program[second]
			}
		case 6:
			if program[first] == 0 {
				jumped = true
				i = program[second]
			}
		case 7:
			if program[first] < program[second] {
				program[third] = 1
			} else {
				program[third] = 0
			}
		case 8:
			if program[first] == program[second] {
				program[third] = 1
			} else {
				program[third] = 0
			}
		case 99:
			return program[0]
		default:
			panic(fmt.Sprintf("invalid opcode %d (program state = %v)", op, program))
		}

		switch op := instr % 100; op {
		case 1, 2, 7, 8:
			i += 4
		case 3, 4:
			i += 2
		case 5, 6:
			if !jumped {
				i += 3
			}
		}
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
	execute(program)
}
