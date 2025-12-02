package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	var left, right []int
	for s.Scan() {
		line := string(s.Bytes())
		space := strings.Index(line, " ")
		if space == -1 {
			log.Fatalf("could not parse line: %#v", line)
		}
		l, err := strconv.Atoi(line[:space])
		if err != nil {
			log.Fatalf("could not parse left of line: %#v: %#v", line, err)
		}
		r, err := strconv.Atoi(line[space+3:])
		if err != nil {
			log.Fatalf("could not parse right of line: %#v: %#v", line, err)
		}
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for n, l := range left {
		sum += int(math.Abs(float64(l - right[n])))
	}

	fmt.Println("Answer:", sum)
}
