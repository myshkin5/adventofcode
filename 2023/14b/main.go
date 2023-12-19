package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"moqueries.org/runtime/hash"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	cliCycleStart := flag.Int("cycle-start", -1, "Iteration cycle starts on")
	cliCycleLength := flag.Int("cycle-length", -1, "Iterations per cycle")
	cliCycleHash := flag.Uint64("cycle-hash", 0, "Hash value of each cycle")
	flag.Parse()
	if *cliCycleStart != -1 || *cliCycleLength != -1 {
		if *cliCycleStart == -1 || *cliCycleLength == -1 {
			log.Fatalf("All or none of the cycle options must be defined")
		}
	}

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

	score := func(pr bool) int {
		all := 0

		for n := 0; n < len(grid); n++ {
			for m := 0; m < len(grid[n]); m++ {
				if pr {
					fmt.Print(string(grid[n][m]))
				}
				if grid[n][m] == 'O' {
					all += len(grid) - n
				}
			}
			if pr {
				fmt.Println()
			}
		}

		return all
	}

	checkCycles := 1000000000
	extraCycles := 0
	if cliCycleStart != nil {
		extraCycles = (checkCycles-*cliCycleStart)%*cliCycleLength - 1
	}

	lastCycle := -1
	var cycleHash hash.Hash
	hashes := map[hash.Hash]struct{}{}
	firstMatch := -1
	for n := 0; n < checkCycles; n++ {
		if n%100000 == 0 {
			log.Println("Cycle:", n, score(false))
		}
		shiftNorth(grid)
		shiftWest(grid)
		shiftSouth(grid)
		shiftEast(grid)

		h := hash.DeepHash(grid)
		if lastCycle != -1 {
			if cycleHash == h {
				diff := n - lastCycle
				lastCycle = n
				log.Println("Cycle found at iteration count:", diff)
			}
		} else {
			if _, ok := hashes[h]; ok {
				lastCycle = n
				cycleHash = h
				log.Println("First cycle found at iteration:", n, "hash:", cycleHash)
			} else {
				hashes[h] = struct{}{}
			}
		}

		if firstMatch != -1 {
			if n-firstMatch == extraCycles {
				break
			}
		}
		if cliCycleHash != nil && hash.Hash(*cliCycleHash) == h && firstMatch == -1 {
			firstMatch = n
		}
	}

	fmt.Println("Answer:", score(true))
}

func shiftNorth(grid [][]rune) {
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
}

func shiftWest(grid [][]rune) {
	for m := 1; m < len(grid[0]); m++ {
		for n := 0; n < len(grid); n++ {
			if grid[n][m] == 'O' {
				mline := -1
				for o := m - 1; o >= 0; o-- {
					if grid[n][o] != '.' {
						break
					}
					mline = o
				}
				if mline != -1 {
					grid[n][m] = '.'
					grid[n][mline] = 'O'
				}
			}
		}
	}
}

func shiftSouth(grid [][]rune) {
	for n := len(grid) - 2; n >= 0; n-- {
		for m := 0; m < len(grid[0]); m++ {
			if grid[n][m] == 'O' {
				mline := -1
				for o := n + 1; o < len(grid); o++ {
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
}

func shiftEast(grid [][]rune) {
	for m := len(grid[0]) - 2; m >= 0; m-- {
		for n := 0; n < len(grid); n++ {
			if grid[n][m] == 'O' {
				mline := -1
				for o := m + 1; o < len(grid[0]); o++ {
					if grid[n][o] != '.' {
						break
					}
					mline = o
				}
				if mline != -1 {
					grid[n][m] = '.'
					grid[n][mline] = 'O'
				}
			}
		}
	}
}
