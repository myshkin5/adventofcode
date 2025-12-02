package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var grid [][]rune
	for s.Scan() {
		line := string(s.Bytes())

		runes := []rune(line)
		grid = append(grid, runes)
	}

	type loc struct {
		x, y int
	}

	var region map[loc]int
	var scan func(r rune, x, y int)
	scan = func(r rune, x, y int) {
		l := loc{x, y}
		if _, ok := region[l]; ok {
			return
		}
		edges := 0
		if x == 0 || grid[y][x-1] != r {
			edges++
		}
		if x == len(grid[y])-1 || grid[y][x+1] != r {
			edges++
		}
		if y == 0 || grid[y-1][x] != r {
			edges++
		}
		if y == len(grid)-1 || grid[y+1][x] != r {
			edges++
		}
		region[l] = edges

		if x > 0 && grid[y][x-1] == r {
			scan(r, x-1, y)
		}
		if x+1 < len(grid[y]) && grid[y][x+1] == r {
			scan(r, x+1, y)
		}
		if y > 0 && grid[y-1][x] == r {
			scan(r, x, y-1)
		}
		if y+1 < len(grid[y]) && grid[y+1][x] == r {
			scan(r, x, y+1)
		}
	}

	sum := 0
	for y, g := range grid {
		for x, r := range g {
			if r == 0 {
				continue
			}
			region = make(map[loc]int)
			scan(r, x, y)

			edges := 0
			for l, e := range region {
				edges += e
				grid[l.y][l.x] = 0
			}
			sum += edges * len(region)
		}
	}

	fmt.Println("Answer:", sum)
}
