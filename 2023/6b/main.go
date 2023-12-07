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
	time := parseLine(string(s.Bytes()), "Time")
	if !s.Scan() {
		log.Fatalf("no record line found in input")
	}
	record := parseLine(string(s.Bytes()), "Distance")

	wins := 0
	for m := 1; m < time; m++ {
		travelTime := time - m
		dist := m * travelTime
		if dist > record {
			wins++
		}
	}

	fmt.Println("Answer:", wins)
}

func parseLine(line, label string) int {
	s1 := strings.Split(line, ": ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ': ': %s", line)
	}
	if s1[0] != label {
		log.Fatalf("didn't find label %s: %s", label, line)
	}

	return strs.Atoi(strings.ReplaceAll(s1[1], " ", ""))
}
