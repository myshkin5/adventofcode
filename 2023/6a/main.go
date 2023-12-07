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
	if !s.Scan() {
		log.Fatalf("no time line found in input")
	}
	times := parseLine(string(s.Bytes()), "Time")
	if !s.Scan() {
		log.Fatalf("no record line found in input")
	}
	records := parseLine(string(s.Bytes()), "Distance")
	if len(times) != len(records) {
		log.Fatalf("mismatched count: %d vs. %d", len(times), len(records))
	}

	factor := 0
	for n := 0; n < len(times); n++ {
		time := times[n]
		record := records[n]

		wins := 0
		for m := 1; m < time; m++ {
			travelTime := time - m
			dist := m * travelTime
			if dist > record {
				wins++
			}
		}

		if factor == 0 {
			factor = wins
		} else {
			factor *= wins
		}
	}

	fmt.Println("Answer:", factor)
}

func parseLine(line, label string) []int {
	s1 := strings.Split(line, ": ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ': ': %s", line)
	}
	if s1[0] != label {
		log.Fatalf("didn't find label %s: %s", label, line)
	}

	s2 := strings.Split(s1[1], " ")
	var vals []int
	for _, val := range s2 {
		if val != "" {
			vals = append(vals, strs.Atoi(val))
		}
	}
	return vals
}
