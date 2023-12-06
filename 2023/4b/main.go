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
	counts := map[int]int{}
	for s.Scan() {
		line := string(s.Bytes())
		checkAll(line, counts)
	}

	sum := 0
	for k, v := range counts {
		fmt.Println(k, v)
		sum += v
	}

	fmt.Println("Answer:", sum)
}

func checkAll(line string, counts map[int]int) {
	s1 := strings.Split(line, ":")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ':': %s", line)
	}
	cId := s1[0]
	nums := s1[1]

	s15 := strings.Split(cId, " ")
	if len(s15) < 2 {
		log.Fatalf("not exactly two sections (%d) at ' ' (card id): %s", len(s15), line)
	}
	cNum := strs.Atoi(s15[len(s15)-1])

	s2 := strings.Split(nums, "|")
	if len(s2) != 2 {
		log.Fatalf("not exactly two sections at '|': %s", line)
	}

	current := counts[cNum] + 1
	counts[cNum] = current
	fmt.Println(cNum, "itself", counts[cNum])
	winners := indexNums(s2[0])
	actuals := indexNums(s2[1])
	val := 0
	for n := range winners {
		if _, ok := actuals[n]; ok {
			val++
			counts[cNum+val] = counts[cNum+val] + current
			fmt.Println(cNum+val, cNum, counts[cNum+val])
		}
	}
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
