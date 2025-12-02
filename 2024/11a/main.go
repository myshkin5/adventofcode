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

	for s.Scan() {
		line := string(s.Bytes())

		splits := strings.Split(line, " ")
		var vals []int
		for _, s := range splits {
			v, err := strconv.Atoi(s)
			if err != nil {
				log.Fatalf("could not parse %#v: %v", s, err)
			}
			vals = append(vals, v)
		}
		fmt.Println(vals)

		for b := 0; b < 25; b++ {
			for n := 0; n < len(vals); n++ {
				if vals[n] == 0 {
					vals[n] = 1
					continue
				}
				digits := int(math.Log10(float64(vals[n]))) + 1
				if digits%2 == 0 {
					pow := int(math.Pow(10, float64(digits/2)))
					first := vals[n] / pow
					second := vals[n] % pow
					vals[n] = first
					vals = slices.Insert(vals, n+1, second)
					n++
					continue
				}
				vals[n] *= 2024
			}
			fmt.Println(vals)
		}

		fmt.Println("Answer:", len(vals))
	}
}
