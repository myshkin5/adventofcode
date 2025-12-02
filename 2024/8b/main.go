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

	type loc struct{ x, y int }
	locs := make(map[rune][]loc)

	var grid [][]rune
	y := 0
	for s.Scan() {
		line := string(s.Bytes())

		runes := []rune(line)
		grid = append(grid, runes)

		for x, r := range runes {
			if r != '.' {
				locs[r] = append(locs[r], loc{x, y})
			}
		}
		y++
	}

	antinodes := func(l loc, dx, dy int) {
		x := l.x
		y := l.y

		for {
			x += dx
			y += dy

			if y < 0 || y >= len(grid) {
				return
			}
			gl := grid[y]
			if x < 0 || x >= len(gl) {
				return
			}
			gl[x] = '#'
		}
	}

	for _, rlocs := range locs {
		for n, loc1 := range rlocs {
			for _, loc2 := range rlocs[n+1:] {
				dx := loc1.x - loc2.x
				dy := loc1.y - loc2.y
				antinodes(loc1, dx, dy)
				antinodes(loc2, -dx, -dy)
			}
		}
	}

	sum := 0
	for _, v := range grid {
		for _, g := range v {
			if g != '.' {
				sum++
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
	fmt.Println("Answer:", sum)
}
