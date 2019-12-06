package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minArr(a []int) int {
	if len(a) == 0 {
		return 0
	}
	res := a[0]
	for i := 1; i < len(a); i += 1 {
		if a[i] < res {
			res = a[i]
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return (-1) * x
	}
	return x
}

type coordinate struct {
	X int
	Y int
}

func dist(c coordinate) int {
	return abs(c.X) + abs(c.Y)
}

type segment struct {
	P1 coordinate
	P2 coordinate
}

func length(s segment) int {
	if vertical(s) {
		return abs(s.P1.Y - s.P2.Y)
	}
	return abs(s.P1.X - s.P2.X)
}

func vertical(s segment) bool {
	return s.P1.X == s.P2.X
}

func between(a, b, x int) bool {
	if a > b {
		return between(b, a, x)
	}
	return x >= a && x <= b
}

func segmentIntersection(s1, s2 segment) (coordinate, bool) {
	if vertical(s1) == vertical(s2) {
		return coordinate{}, false
	}
	if !vertical(s1) {
		return segmentIntersection(s2, s1)
	}
	// s1 is vertical, s2 is horizontal
	if between(s2.P1.X, s2.P2.X, s1.P1.X) && between(s1.P1.Y, s1.P2.Y, s2.P1.Y) {
		return coordinate{s1.P1.X, s2.P1.Y}, true
	}
	return coordinate{}, false

}

func lineIntersections(l1, l2 []segment) []coordinate {
	res := []coordinate{}
	for _, s1 := range l1 {
		for _, s2 := range l2 {
			c, ok := segmentIntersection(s1, s2)
			if ok && (c != coordinate{0, 0}) {
				res = append(res, c)
			}
		}
	}
	return res
}

func parseSegments(s string) []segment {
	res := []segment{}
	current := coordinate{0, 0}
	for _, step := range strings.Split(s, ",") {
		dir := step[0]
		dist, err := strconv.Atoi(step[1:])
		if err != nil {
			panic(err)
		}
		switch dir {
		case 'L':
			res = append(res, segment{current, coordinate{current.X - dist, current.Y}})
		case 'R':
			res = append(res, segment{current, coordinate{current.X + dist, current.Y}})
		case 'U':
			res = append(res, segment{current, coordinate{current.X, current.Y + dist}})
		case 'D':
			res = append(res, segment{current, coordinate{current.X, current.Y - dist}})
		}
		current = res[len(res)-1].P2
	}
	return res
}

func contains(s segment, c coordinate) (int, bool) {
	if vertical(s) {
		if s.P1.X == c.X && between(s.P1.Y, s.P2.Y, c.Y) {
			return abs(s.P1.Y - c.Y), true
		}
	} else {
		if s.P1.Y == c.Y && between(s.P1.X, s.P2.X, c.X) {
			return abs(s.P1.X - c.X), true
		}
	}
	return -1, false
}

func search(line []segment, target coordinate) (int, bool) {
	res := 0
	for _, s := range line {
		if d, ok := contains(s, target); ok {
			return res + d, true
		}
		res += length(s)
	}
	return 0, false
}

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(bytes), "\n")

	line1 := parseSegments(parts[0])
	line2 := parseSegments(parts[1])

	res := math.MaxInt32
	for _, i := range lineIntersections(line1, line2) {
		d1, _ := search(line1, i)
		d2, _ := search(line2, i)
		if d1+d2 < res {
			res = d1 + d2
		}
	}

	fmt.Println(res)
}
