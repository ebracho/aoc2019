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

func ancestors(o *object) map[string]int {
	res := map[string]int{}
	hops := 0
	for ; o != nil; o = o.orbiting {
		res[o.name] = hops
		hops += 1
	}
	return res
}

func commonAncestor(x, y *object) (*object, int, int) {
	xAncestors := ancestors(x)
	yHops := 0
	for ; y != nil; y = y.orbiting {
		if xHops, ok := xAncestors[y.name]; ok {
			return y, xHops, yHops
		}
		yHops += 1
	}
	return nil, 0, 0
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
	_, youHops, sanHops := commonAncestor(u.get("YOU"), u.get("SAN"))
	fmt.Println(youHops + sanHops - 2)
}
