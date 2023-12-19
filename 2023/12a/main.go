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

	all := 0

	for s.Scan() {
		total := 0

		line := string(s.Bytes())

		s1 := strings.Split(line, " ")
		if len(s1) != 2 {
			log.Fatalf("not two elementes at ' ': %s", line)
		}
		reading := s1[0]

		s2 := strings.Split(s1[1], ",")
		var counts []int
		for _, s := range s2 {
			counts = append(counts, strs.Atoi(s))
		}

		var stack []int
		cur := -1
		springStart := -1
		popStack := func() bool {
			if len(stack) == 0 {
				fmt.Println(line, total)
				fmt.Println()
				all += total
				return true
			}
			cur = stack[len(stack)-1]
			springStart = -1
			stack = stack[:len(stack)-1]
			return false
		}
		for {
			cur++
			pastKnownWell := cur > 0 && cur < len(reading) && reading[cur-1] == '#' && springStart == -1
			if cur >= len(reading) || pastKnownWell {
				if popStack() {
					break
				}
				continue
			}

			if reading[cur] == '.' {
				if springStart != -1 {
					cur = springStart
					springStart = -1
				}
				continue
			}
			if (reading[cur] == '?' || reading[cur] == '#') && springStart == -1 {
				if cur == 0 || reading[cur-1] != '#' {
					springStart = cur
				}
			}
			sLen := counts[len(stack)]
			if springStart != -1 && cur-springStart+1 == sLen {
				if cur == len(reading)-1 || reading[cur+1] == '.' || reading[cur+1] == '?' {
					stack = append(stack, springStart)
					cur++
					springStart = -1
					if len(stack) == len(counts) {
						// if known wells (#) remain, got to pop the stack
						if strs.IndexAt(reading, "#", cur) != -1 {
							if popStack() {
								break
							}
							continue
						}
						fmt.Println(stack)
						total++
						cur -= sLen
						stack = stack[:len(stack)-1]
					}
				} else {
					cur = springStart
					springStart = -1
				}
			}
		}
	}

	fmt.Println("Answer:", all)
}
