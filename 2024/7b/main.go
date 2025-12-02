package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

		oc := int(math.Pow(3, float64(len(nums)-1)))
		for ops := 0; ops < oc; ops++ {
			cumul := nums[0]
			dops := ops
			for _, num := range nums[1:] {
				bit := dops % 3
				if bit == 0 {
					cumul *= num
				} else if bit == 1 {
					cumul += num
				} else {
					cumul = cumul*int(math.Pow(10, float64(digits(num)))) + num
				}
				dops /= 3
			}
			if val == cumul {
				sum += val
				break
			}
		}
	}

	fmt.Println("Answer:", sum)
}

func digits(n int) int {
	if n == 0 {
		return 1
	}
	var sign int
	if n < 0 {
		sign = 1
		n = -n
	}
	return sign + int(math.Floor(math.Log10(float64(n)))) + 1
}
