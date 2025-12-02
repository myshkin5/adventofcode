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
			if grid[n][m] != 'X' {
				continue
			}
			if m+3 < len(grid[n]) && grid[n][m+1] == 'M' && grid[n][m+2] == 'A' && grid[n][m+3] == 'S' {
				// fmt.Println("right", n, m)
				sum++
			}
			if m+3 < len(grid[n]) && n+3 < len(grid) && grid[n+1][m+1] == 'M' && grid[n+2][m+2] == 'A' && grid[n+3][m+3] == 'S' {
				// fmt.Println("down right", n, m)
				sum++
			}
			if n+3 < len(grid) && grid[n+1][m] == 'M' && grid[n+2][m] == 'A' && grid[n+3][m] == 'S' {
				// fmt.Println("down", n, m)
				sum++
			}
			if m-3 >= 0 && n+3 < len(grid) && grid[n+1][m-1] == 'M' && grid[n+2][m-2] == 'A' && grid[n+3][m-3] == 'S' {
				// fmt.Println("down left", n, m)
				sum++
			}
			if m-3 >= 0 && grid[n][m-1] == 'M' && grid[n][m-2] == 'A' && grid[n][m-3] == 'S' {
				// fmt.Println("left", n, m)
				sum++
			}
			if m-3 >= 0 && n-3 >= 0 && grid[n-1][m-1] == 'M' && grid[n-2][m-2] == 'A' && grid[n-3][m-3] == 'S' {
				// fmt.Println("up left", n, m)
				sum++
			}
			if n-3 >= 0 && grid[n-1][m] == 'M' && grid[n-2][m] == 'A' && grid[n-3][m] == 'S' {
				// fmt.Println("up", n, m)
				sum++
			}
			if m+3 < len(grid[n]) && n-3 >= 0 && grid[n-1][m+1] == 'M' && grid[n-2][m+2] == 'A' && grid[n-3][m+3] == 'S' {
				// fmt.Println("up right", n, m)
				sum++
			}
		}
	}

	fmt.Println("Answer:", sum)
}
