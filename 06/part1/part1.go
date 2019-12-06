package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type object struct {
	name     string
	orbiting *object
}

type universe map[string]*object

func (u universe) get(name string) *object {
	_, ok := u[name]
	if !ok {
		u[name] = &object{name: name}
	}
	return u[name]
}

func (u universe) orbits() int {
	subOrbitsCache := map[string]int{}
	var subOrbits func(obj *object) int
	subOrbits = func(obj *object) (res int) {
		if r, ok := subOrbitsCache[obj.name]; ok {
			return r
		}
		defer func() {
			subOrbitsCache[obj.name] = res
		}()
		if obj.orbiting != nil {
			return subOrbits(obj.orbiting) + 1
		}
		return 0
	}
	res := 0
	for _, o := range u {
		res += subOrbits(o)
	}
	return res
}

func main() {
	in, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	u := universe{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ")")
		parent := u.get(parts[0])
		child := u.get(parts[1])
		child.orbiting = parent
	}
	if scanner.Err() != nil {
		panic(err)
	}
	fmt.Println(u.orbits())
}
