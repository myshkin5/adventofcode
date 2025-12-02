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

type direction int

const (
	up direction = 1 << iota
	right
	down
	left
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var grid [][]rune
	x, y, c := -1, -1, 0
	for s.Scan() {
		line := string(s.Bytes())

		if start := strings.Index(line, "^"); start != -1 {
			x = start
			y = c
		}

		grid = append(grid, []rune(line))
		c++
	}

	fmt.Printf("Starting at %d, %d\n", x, y)

	sum := 1
	dir := up
	for {
		out := true
		switch dir {
		case up:
			for ny := y - 1; ny >= 0; ny-- {
				if grid[ny][x] == '.' {
					sum++
					grid[ny][x] = 'X'
				}
				if grid[ny][x] == '#' {
					y = ny + 1
					dir = right
					out = false
					break
				}
			}
		case right:
			for nx := x + 1; nx < len(grid[y]); nx++ {
				if grid[y][nx] == '.' {
					sum++
					grid[y][nx] = 'X'
				}
				if grid[y][nx] == '#' {
					x = nx - 1
					dir = down
					out = false
					break
				}
			}
		case down:
			for ny := y + 1; ny < len(grid); ny++ {
				if grid[ny][x] == '.' {
					sum++
					grid[ny][x] = 'X'
				}
				if grid[ny][x] == '#' {
					y = ny - 1
					dir = left
					out = false
					break
				}
			}
		case left:
			for nx := x - 1; nx >= 0; nx-- {
				if grid[y][nx] == '.' {
					sum++
					grid[y][nx] = 'X'
				}
				if grid[y][nx] == '#' {
					x = nx + 1
					dir = up
					out = false
					break
				}
			}
		}

		for _, g := range grid {
			for _, r := range g {
				fmt.Print(string(r))
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()

		if out {
			break
		}
	}

	fmt.Println("Answer:", sum)
}
