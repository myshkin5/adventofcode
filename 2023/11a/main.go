package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	xlen := 0
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

		grid = append(grid, line)
		if !strings.Contains(line, "#") {
			// h expansion
			grid = append(grid, line)
		}
	}

	for n := len(grid[0]) - 1; n >= 0; n-- {
		allDots := true
		for m := 0; m < len(grid); m++ {
			if grid[m][n] != '.' {
				allDots = false
				break
			}
		}

		if allDots {
			for m := 0; m < len(grid); m++ {
				grid[m] = grid[m][0:n] + "." + grid[m][n:]
			}
		}
	}

	var gs []loc
	for n := 0; n < len(grid); n++ {
		fmt.Println(grid[n])
		m := -1
		for {
			m = strs.IndexAt(grid[n], "#", m+1)
			if m == -1 {
				break
			}
			gs = append(gs, loc{x: m, y: n})
		}
	}

	fmt.Println("Galaxies:", len(gs))

	total := 0
	for n := 0; n < len(gs); n++ {
		for m := n + 1; m < len(gs); m++ {
			total += int(math.Abs(float64(gs[m].x - gs[n].x)))
			total += gs[m].y - gs[n].y
		}
	}

	fmt.Println("Answer:", total)
}

type loc struct {
	x, y int
}
