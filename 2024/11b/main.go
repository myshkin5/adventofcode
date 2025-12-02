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

	type pow struct {
		min, max, pow uint64
	}
	pows := []pow{{min: 10, max: 99, pow: 10}}
	minV := uint64(10)
	powV := uint64(10)
	for n := 0; n < 10; n++ {
		minV *= 100
		powV *= 10
		pows = append(pows, pow{min: minV, max: minV*10 - 1, pow: powV})
	}

	type answerKey struct {
		blink int
		val   uint64
	}
	answers := make(map[answerKey]int)

	sum := 0
	for s.Scan() {
		line := string(s.Bytes())

		splits := strings.Split(line, " ")
		for _, s := range splits {
			v, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				log.Fatalf("could not parse %#v: %v", s, err)
			}

			fmt.Println(v)

			var split func(b int, v uint64) int
			split = func(b int, v uint64) int {
				if b == 75 {
					// fmt.Println("end  ", b, v, 1)
					return 1
				}

				k := answerKey{blink: b, val: v}
				if r, ok := answers[k]; ok {
					// fmt.Println("ans  ", b, v, r)
					return r
				}

				if v == 0 {
					r := split(b+1, 1)
					answers[k] = r
					// fmt.Println("zero ", b, v, r)
					return r
				}
				if v > pows[len(pows)-1].max {
					log.Panicln(v)
				}
				for _, p := range pows {
					if v < p.min {
						break
					}
					if v <= p.max {
						first := split(b+1, v/p.pow)
						second := split(b+1, v%p.pow)
						r := first + second
						answers[k] = r
						// fmt.Println("split", b, v, r, first, second)
						return r
					}
				}
				r := split(b+1, v*2024)
				answers[k] = r
				// fmt.Println("year ", b, v, r)
				return r
			}

			sum += split(0, uint64(v))
			fmt.Println()
			fmt.Println()
		}

		fmt.Println("Answer:", sum)
	}
}
