package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/myshkin5/adventofcode/paths"
)

type direction int

const (
	up direction = 1 << iota
	right
	down
	left
)

type path struct {
	dir  direction
	x, y int
}

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var grid []string
	for s.Scan() {
		line := string(s.Bytes())

		grid = append(grid, line)
	}

	egrid := make([][]direction, len(grid))
	for n := 0; n < len(grid); n++ {
		egrid[n] = make([]direction, len(grid[0]))
	}

	allPaths := []path{{dir: right, x: -1, y: 0}}
	for n := 0; n < len(allPaths); n++ {
		cur := allPaths[n]
		for {
			x, y := move(cur)

			// Done with this path if we are off the grid
			if x < 0 || x >= len(grid[0]) {
				break
			}
			if y < 0 || y >= len(grid) {
				break
			}

			if egrid[y][x]&cur.dir != 0 {
				// Done with this path if we have already been here (don't loop)
				break
			}
			egrid[y][x] |= cur.dir
			switch r := grid[y][x]; r {
			case '.':
				// continue
			case '|':
				allPaths = append(allPaths, path{dir: up, x: x, y: y})
				cur.dir = down
			case '-':
				allPaths = append(allPaths, path{dir: left, x: x, y: y})
				cur.dir = right
			case '/':
				switch cur.dir {
				case up:
					cur.dir = right
				case right:
					cur.dir = up
				case down:
					cur.dir = left
				case left:
					cur.dir = down
				}
			case '\\':
				switch cur.dir {
				case up:
					cur.dir = left
				case right:
					cur.dir = down
				case down:
					cur.dir = right
				case left:
					cur.dir = up
				}
			default:
				log.Fatalln("Unknown token:", string(r), "x:", cur.x, "y:", cur.y)
			}

			cur.x = x
			cur.y = y
		}
	}

	all := 0

	for n := 0; n < len(egrid); n++ {
		for m := 0; m < len(egrid[n]); m++ {
			if egrid[n][m] != 0 {
				all++
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	fmt.Println("Answer:", all)
}

func move(p path) (int, int) {
	switch p.dir {
	case right:
		return p.x + 1, p.y
	case left:
		return p.x - 1, p.y
	case up:
		return p.x, p.y - 1
	case down:
		return p.x, p.y + 1
	default:
		log.Fatalf("Unknown direction: %d", p.dir)
	}
	return -1, -1
}
