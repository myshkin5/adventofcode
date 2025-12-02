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

	var grid [][]int
	for s.Scan() {
		line := string(s.Bytes())

		var vals []int
		for _, v := range line {
			vals = append(vals, int(v-'0'))
		}
		grid = append(grid, vals)
	}

	type loc struct{ x, y int }
	nines := make(map[loc]struct{})
	var search func(x, y, val int) int
	search = func(x, y, val int) int {
		if val == 10 {
			l := loc{x, y}
			if _, ok := nines[l]; !ok {
				fmt.Printf("Found: %d, %d\n", x, y)
				nines[l] = struct{}{}
				return 1
			}
			return 0
		}
		count := 0
		if x > 0 && grid[y][x-1] == val {
			count += search(x-1, y, val+1)
		}
		if x < len(grid[y])-1 && grid[y][x+1] == val {
			count += search(x+1, y, val+1)
		}
		if y > 0 && grid[y-1][x] == val {
			count += search(x, y-1, val+1)
		}
		if y < len(grid)-1 && grid[y+1][x] == val {
			count += search(x, y+1, val+1)
		}
		return count
	}

	sum := 0
	for y, vals := range grid {
		for x, val := range vals {
			if val == 0 {
				clear(nines)
				fmt.Printf("Trail head %d, %d\n", x, y)
				th := search(x, y, 1)
				fmt.Printf("Score: %d\n\n\n", th)
				sum += th
			}
		}
	}

	for _, g := range grid {
		for _, r := range g {
			fmt.Print(r)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Answer:", sum)
}
