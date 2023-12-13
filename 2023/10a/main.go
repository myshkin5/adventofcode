package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	xlen := 0
	start := loc{x: -1, y: -1}
	var grid []string
	for s.Scan() {
		line := string(s.Bytes())
		if xlen == 0 {
			xlen = len(line)
		} else {
			if len(line) != xlen {
				log.Fatalf("X length so far %d doesn't match length of %s", xlen, line)
			}
		}

		s := strings.Index(line, "S")
		if s != -1 {
			if start.x != -1 || start.y != -1 || strings.LastIndex(line, "S") != s {
				log.Fatalf("found more than one start, first: %#v, current: %d, %d",
					start, s, len(grid))
			}
			start.x = s
			start.y = len(grid)
		}

		grid = append(grid, line)
	}

	if start.x == -1 || start.y == -1 {
		log.Fatalf("did not find start")
	}

	log.Println("Start:", start)

	nextA, nextB := findAllConnections(start, grid)
	curA, curB := start, start
	total := 0
	for n := 0; n < 10000000; n++ {
		curA = curA.move(nextA)
		curB = curB.move(nextB)
		log.Println("Step:", n, "cur a:", curA, "cur b:", curB)

		var ok bool
		ok, nextA = next(curA, nextA, grid)
		if !ok {
			log.Fatalf("location %#v is not connected when coming from %d", curA, nextA)
		}
		ok, nextB = next(curB, nextB, grid)
		if !ok {
			log.Fatalf("location %#v is not connected when coming from %d", curB, nextB)
		}

		if curA == curB {
			total = n
			break
		}
	}

	fmt.Println("Answer:", total+1)
}

type loc struct {
	x, y int
}

func (l loc) isNeighbor(neighbor loc) bool {
	if l == neighbor {
		return false
	}
	if l.x < neighbor.x-1 || l.x > neighbor.x+1 || l.y < neighbor.y-1 || l.y > neighbor.y+1 {
		return false
	}

	return true
}

type direction int

const (
	up direction = iota
	upRight
	right
	downRight
	down
	downLeft
	left
	upLeft
)

func (l loc) move(dir direction) loc {
	switch dir {
	case up:
		return loc{x: l.x, y: l.y - 1}
	case upRight:
		return loc{x: l.x + 1, y: l.y - 1}
	case right:
		return loc{x: l.x + 1, y: l.y}
	case downRight:
		return loc{x: l.x + 1, y: l.y + 1}
	case down:
		return loc{x: l.x, y: l.y + 1}
	case downLeft:
		return loc{x: l.x - 1, y: l.y + 1}
	case left:
		return loc{x: l.x - 1, y: l.y}
	case upLeft:
		return loc{x: l.x - 1, y: l.y - 1}
	default:
		log.Fatalf("unknown direction %d", dir)
	}

	return loc{}
}

var connected = map[rune]map[direction]direction{
	'7': {right: down, up: left},
	'J': {down: left, right: up},
	'L': {left: up, down: right},
	'F': {up: right, left: down},
	'|': {up: up, down: down},
	'-': {left: left, right: right},
	'.': {},
}

func findAllConnections(start loc, grid []string) (direction, direction) {
	var dirs []direction
	for d := up; d <= upLeft; d++ {
		nloc := start.move(d)
		cOk, _ := next(nloc, d, grid)
		if cOk {
			dirs = append(dirs, d)
		}
	}

	if len(dirs) != 2 {
		log.Fatalf("start not connected to exactly two neighbors")
	}

	return dirs[0], dirs[1]
}

func next(l loc, d direction, grid []string) (bool, direction) {
	if l.y < 0 || l.y >= len(grid) || l.x < 0 || l.x >= len(grid[0]) {
		return false, 0
	}

	nr := grid[l.y][l.x]
	cdirs := connected[rune(nr)]
	if nd, ok := cdirs[d]; ok {
		return true, nd
	}

	return false, 0
}
