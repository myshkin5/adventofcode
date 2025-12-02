package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	sum := 0
	for s.Scan() {
		line := string(s.Bytes())

		splits := strings.Split(line, ": ")
		if len(splits) != 2 {
			log.Fatalf("invalid line: %#v", line)
		}
		val, err := strconv.Atoi(splits[0])
		if err != nil {
			log.Fatalf("invalid line: %#v", line)
		}

		splits = strings.Split(splits[1], " ")
		var nums []int
		for _, s := range splits {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("invalid line: %#v", line)
			}
			nums = append(nums, num)
		}

		for ops := 0; ops < 1<<(len(nums)-1); ops++ {
			cumul := nums[0]
			for n, num := range nums[1:] {
				bit := (1 << n) & ops
				if bit == 0 {
					cumul *= num
				} else {
					cumul += num
				}
			}
			if val == cumul {
				sum += val
				break
			}
		}
	}

	fmt.Println("Answer:", sum)
}
