package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/myshkin5/adventofcode/paths"
)

var rexp = regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|[0-9])")
var sToI = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	sum := 0
	for s.Scan() {
		line := string(s.Bytes())
		foundFirst := false
		var first, last int
		for n := 0; n < len(line); n++ {
			word := rexp.FindString(line[n:])
			if word == "" {
				break
			}
			var ok bool
			last, ok = sToI[word]
			if !ok {
				last, err = strconv.Atoi(word)
				if err != nil {
					log.Fatalf("non-digit found on line: %s", line)
				}
			}

			if !foundFirst {
				foundFirst = true
				first = last
			}
		}
		sum += first*10 + last
		fmt.Println(first*10+last, line)
	}

	fmt.Println("Answer:", sum)
}
