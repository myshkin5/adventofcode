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

	var grid []string
	for s.Scan() {
		line := string(s.Bytes())

		grid = append(grid, line)
	}

	sum := 0
	for n := 0; n < len(grid); n++ {
		for m := 0; m < len(grid[n]); m++ {
			if grid[n][m] != 'M' {
				continue
			}
			if m+2 < len(grid[n]) && n+2 < len(grid) &&
				grid[n+1][m+1] == 'A' && grid[n+2][m+2] == 'S' &&
				grid[n+2][m] == 'M' && grid[n][m+2] == 'S' {
				sum++
			}
			if m+2 < len(grid[n]) && n+2 < len(grid) &&
				grid[n+1][m+1] == 'A' && grid[n+2][m+2] == 'S' &&
				grid[n+2][m] == 'S' && grid[n][m+2] == 'M' {
				sum++
			}
			if m-2 >= 0 && n-2 >= 0 &&
				grid[n-1][m-1] == 'A' && grid[n-2][m-2] == 'S' &&
				grid[n-2][m] == 'M' && grid[n][m-2] == 'S' {
				sum++
			}
			if m-2 >= 0 && n-2 >= 0 &&
				grid[n-1][m-1] == 'A' && grid[n-2][m-2] == 'S' &&
				grid[n-2][m] == 'S' && grid[n][m-2] == 'M' {
				sum++
			}
		}
	}

	fmt.Println("Answer:", sum)
}
