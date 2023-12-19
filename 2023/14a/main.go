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

		rs := make([]rune, len(line))
		for n, r := range line {
			rs[n] = r
		}
		grid = append(grid, rs)
	}

	for n := 1; n < len(grid); n++ {
		for m := 0; m < len(grid[0]); m++ {
			if grid[n][m] == 'O' {
				mline := -1
				for o := n - 1; o >= 0; o-- {
					if grid[o][m] != '.' {
						break
					}
					mline = o
				}
				if mline != -1 {
					grid[n][m] = '.'
					grid[mline][m] = 'O'
				}
			}
		}
	}

	all := 0

	for n := 0; n < len(grid); n++ {
		for m := 0; m < len(grid[n]); m++ {
			fmt.Print(string(grid[n][m]))
			if grid[n][m] == 'O' {
				all += len(grid) - n
			}
		}
		fmt.Println()
	}

	fmt.Println("Answer:", all)
}
