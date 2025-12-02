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
	startx, starty, c := -1, -1, 0
	for s.Scan() {
		line := string(s.Bytes())

		if start := strings.Index(line, "^"); start != -1 {
			startx = start
			starty = c
		}

		grid = append(grid, []rune(line))
		c++
	}

	check := func(mark bool) bool {
		x := startx
		y := starty
		dir := up

		type locDir struct {
			dir  direction
			x, y int
		}
		steps := make(map[locDir]struct{})
		for {
			out := true
			switch dir {
			case up:
				for ny := y - 1; ny >= 0; ny-- {
					if grid[ny][x] == '#' || grid[ny][x] == 'O' {
						y = ny + 1
						dir = right
						out = false
						break
					}

					l := locDir{dir: dir, x: x, y: ny}
					if _, ok := steps[l]; ok {
						return true
					}
					steps[l] = struct{}{}
					if mark {
						grid[ny][x] = 'X'
					}
				}
			case right:
				for nx := x + 1; nx < len(grid[y]); nx++ {
					if grid[y][nx] == '#' || grid[y][nx] == 'O' {
						x = nx - 1
						dir = down
						out = false
						break
					}

					l := locDir{dir: dir, x: nx, y: y}
					if _, ok := steps[l]; ok {
						return true
					}
					steps[l] = struct{}{}
					if mark {
						grid[y][nx] = 'X'
					}
				}
			case down:
				for ny := y + 1; ny < len(grid); ny++ {
					if grid[ny][x] == '#' || grid[ny][x] == 'O' {
						y = ny - 1
						dir = left
						out = false
						break
					}
					l := locDir{dir: dir, x: x, y: ny}
					if _, ok := steps[l]; ok {
						return true
					}
					steps[l] = struct{}{}
					if mark {
						grid[ny][x] = 'X'
					}
				}
			case left:
				for nx := x - 1; nx >= 0; nx-- {
					if grid[y][nx] == '#' || grid[y][nx] == 'O' {
						x = nx + 1
						dir = up
						out = false
						break
					}

					l := locDir{dir: dir, x: nx, y: y}
					if _, ok := steps[l]; ok {
						return true
					}
					steps[l] = struct{}{}
					if mark {
						grid[y][nx] = 'X'
					}
				}
			}

			if out {
				break
			}
		}
		return false
	}
	check(true)
	sum := 0
	for y, v := range grid {
		for x, g := range v {
			if g == 'X' {
				grid[y][x] = 'O'
				if check(false) {
					fmt.Println("Cycle")
					sum++
				}
				// for _, g := range grid {
				// 	for _, r := range g {
				// 		fmt.Print(string(r))
				// 	}
				// 	fmt.Println()
				// }
				// fmt.Println()
				// fmt.Println()
				// fmt.Println(sum)
				grid[y][x] = 'X'
			}
		}
	}

	fmt.Println("Answer:", sum)
}
