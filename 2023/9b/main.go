package main

import (
	"bufio"
	"fmt"
	"log"
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

	total := 0
	for s.Scan() {
		s1 := strings.Split(string(s.Bytes()), " ")
		var readings []int
		for _, reading := range s1 {
			readings = append(readings, strs.Atoi(reading))
		}

		total += extrapolate(readings)
	}

	fmt.Println("Answer:", total)
}

func extrapolate(readings []int) int {
	log.Println("Readings:", readings)
	var levels [][]int
	levels = append(levels, readings)
	next := readings
	for {
		var cur []int
		allZeros := true
		for n := 1; n < len(next); n++ {
			val := next[n] - next[n-1]
			if val != 0 {
				allZeros = false
			}
			cur = append(cur, val)
		}

		if allZeros {
			break
		}
		levels = append(levels, cur)
		next = cur
	}

	same := levels[len(levels)-1]
	val := 0
	for _, s := range same {
		if val == 0 {
			val = s
		} else {
			if val != s {
				log.Fatalf("last line not all the same: %#v", readings)
			}
		}
	}
	log.Println("First:", val, "levels:", len(levels))

	for n := len(levels) - 2; n >= 0; n-- {
		level := levels[n]
		val = level[0] - val
		log.Println("Subsequent:", val)
	}

	return val
}
