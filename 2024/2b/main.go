package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := string(s.Bytes())

		var levels []int
		prev := 0
		for n, c := range line {
			if c == ' ' || n == len(line)-1 {
				last := n
				if n == len(line)-1 {
					last++
				}
				l, err := strconv.Atoi(line[prev:last])
				if err != nil {
					log.Fatalf("could not parse level %d, from %d to %d on line %s: %v", len(levels), prev, n, line, err)
				}
				levels = append(levels, l)
				prev = n + 1
			}
		}

		fmt.Println(levels, len(levels))
		if len(levels) < 2 {
			log.Fatalf("line %s has less than two levels", line)
		}

		check := func(drop int) bool {
			fmt.Println("Dropping", drop)
			prev := -1
			accum := 0
			for n := 0; n < len(levels); n++ {
				if n == drop {
					continue
				}
				if prev == -1 {
					prev = levels[n]
					continue
				}
				diff := levels[n] - prev
				if diff < 0 && accum > 0 {
					fmt.Println("Decrementing after incrementing", n)
					return false
				} else if diff > 0 && accum < 0 {
					fmt.Println("Incrementing after decrementing", n)
					return false
				} else if diff == 0 {
					fmt.Println("No change", n)
					return false
				} else if math.Abs(float64(diff)) > 3 {
					fmt.Println("Leap", n, diff)
					return false
				}
				accum += diff
				prev = levels[n]
			}
			fmt.Println("Good")
			return true
		}

		if check(-1) {
			sum++
			continue
		}
		found := false
		for drop := 0; drop < len(levels); drop++ {
			if check(drop) {
				found = true
				break
			}
		}
		if found {
			sum++
		}
	}

	fmt.Println("Answer:", sum)
}
