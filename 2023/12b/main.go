package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	all := int64(0)
	solutions := int32(0)

	doIt := func(wId int, line string) {
		start := time.Now()

		total := 0

		s1 := strings.Split(line, " ")
		if len(s1) != 2 {
			log.Fatalf("not two elementes at ' ': %s", line)
		}
		origReading := s1[0]
		reading := fmt.Sprintf("%s?%s?%s?%s?%s",
			origReading, origReading, origReading, origReading, origReading)

		s2 := strings.Split(s1[1], ",")
		var counts []int
		for n := 0; n < 5; n++ {
			for _, s := range s2 {
				counts = append(counts, strs.Atoi(s))
			}
		}

		var stack []int
		cur := -1
		springStart := -1
		popStack := func() bool {
			if len(stack) == 0 {
				s := atomic.AddInt32(&solutions, 1)
				log.Println(wId, time.Since(start).String(), s, line, total)

				atomic.AddInt64(&all, int64(total))
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

	work := make(chan string, 2000)
	for s.Scan() {
		work <- string(s.Bytes())
	}
	close(work)

	wg := sync.WaitGroup{}
	for n := 0; n < 8; n++ {
		wg.Add(1)
		go func(wId int) {
			for line := range work {
				doIt(wId, line)
			}
			wg.Done()
		}(n)
	}

	wg.Wait()

	log.Println("Answer:", all)
}
