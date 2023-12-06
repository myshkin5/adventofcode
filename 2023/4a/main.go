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
	var sum int
	for s.Scan() {
		line := string(s.Bytes())
		sum += checkAll(line)
	}

	fmt.Println("Answer:", sum)
}

func checkAll(line string) int {
	s1 := strings.Split(line, ":")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ':': %s", line)
	}
	cId := s1[0]
	nums := s1[1]

	s2 := strings.Split(nums, "|")
	if len(s2) != 2 {
		log.Fatalf("not exactly two sections at '|': %s", line)
	}

	winners := indexNums(s2[0])
	actuals := indexNums(s2[1])
	val := 0
	for n := range winners {
		if _, ok := actuals[n]; ok {
			if val == 0 {
				val = 1
			} else {
				val *= 2
			}
		}
	}
	fmt.Println(cId, val)
	return val
}

func indexNums(nums string) map[int]struct{} {
	ns := map[int]struct{}{}
	for _, s := range strings.Split(nums, " ") {
		if s != "" {
			ns[strs.Atoi(s)] = struct{}{}
		}
	}

	return ns
}
