package main

import (
	"bufio"
	"fmt"
	"log"
	"math/bits"
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

	all := 0

	for pIdx := 0; true; pIdx++ {
		var spattern []string
		for s.Scan() {
			line := string(s.Bytes())

			if line == "" {
				break
			}

			spattern = append(spattern, line)
		}

		if len(spattern) == 0 {
			break
		}

		hpattern := make([]uint, len(spattern))
		vpattern := make([]uint, len(spattern[0]))
		for n, line := range spattern {
			for m, r := range line {
				if r == '#' {
					hpattern[n] |= 1 << m
					vpattern[m] |= 1 << n
				}
			}
		}

		checkMatch := func(a, b uint) (bool, bool) {
			switch bits.OnesCount(a ^ b) {
			case 0:
				return true, false
			case 1:
				return true, true
			default:
				return false, false
			}
		}

		check := func(pattern []uint, multiplier int) int {
			last := uint(0)
			for n, v := range pattern {
				match, smudge := checkMatch(last, v)
				if n > 0 && match {
					symmetry := true
					for m := n + 1; m < len(pattern); m++ {
						b := 2*n - m - 1
						if b < 0 {
							break
						}
						sm, ss := checkMatch(pattern[b], pattern[m])
						if !sm || (ss && smudge) {
							symmetry = false
							break
						}
						if ss {
							smudge = true
						}
					}
					if symmetry && smudge {
						all += n * multiplier
						return n - 1
					}
				}
				last = v
			}
			return -1
		}

		hfound, vfound := -1, -1
		hfound = check(hpattern, 100)
		if hfound == -1 {
			vfound = check(vpattern, 1)
		}
		fmt.Println("Pattern:", pIdx, hpattern, hfound, vpattern, vfound)
	}

	fmt.Println("Answer:", all)
}
