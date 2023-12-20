package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var line string
	for s.Scan() {
		if line != "" {
			log.Fatalf("more than one line")
		}
		line = string(s.Bytes())
	}

	total := 0
	for _, step := range strings.Split(line, ",") {
		total += hash(step)
	}

	fmt.Println("Answer:", total)
}

func hash(step string) int {
	val := 0
	for _, r := range step {
		val += int(r)
		val *= 17
		val %= 256
	}
	// fmt.Println("Step:", step, "hash:", val)
	return val
}
