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
	var hEmpties []int
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
			hEmpties = append(hEmpties, len(grid)-1)
			fmt.Println("H:", len(grid)-1)
		}
	}

	var vEmpties []int
	for n := 0; n < len(grid[0]); n++ {
		allDots := true
		for m := 0; m < len(grid); m++ {
			if grid[m][n] != '.' {
				allDots = false
				break
			}
		}

		if allDots {
			vEmpties = append(vEmpties, n)
			fmt.Println("V:", n)
		}
	}

	var gs []loc
	for n := 0; n < len(grid); n++ {
		m := -1
		for {
			m = strs.IndexAt(grid[n], "#", m+1)
			if m == -1 {
				break
			}
			gs = append(gs, loc{x: m, y: n})
		}
	}

	for n := 0; n < len(gs); n++ {
		fmt.Println("Galaxy:", n, gs[n])
	}
	fmt.Println("Galaxies:", len(gs))

	xp := 999999
	for n := range hEmpties {
		for m := n + 1; m < len(hEmpties); m++ {
			hEmpties[m] += xp
		}
		for m := range gs {
			if gs[m].y > hEmpties[n] {
				gs[m].y += xp
			}
		}
	}

	for n := range vEmpties {
		for m := n + 1; m < len(vEmpties); m++ {
			vEmpties[m] += xp
		}
		for m := range gs {
			if gs[m].x > vEmpties[n] {
				gs[m].x += xp
			}
		}
	}

	for n := 0; n < len(gs); n++ {
		fmt.Println("Galaxy:", n, gs[n])
	}

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
