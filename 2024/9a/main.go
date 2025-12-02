package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

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

		var disk []int
		id := 0
		for n := 0; n < len(line); n += 2 {
			count, err := strconv.Atoi(line[n : n+1])
			if err != nil {
				log.Fatalf("could not parse %q at %d: %v", line[n:n+1], n, err)
			}
			for i := 0; i < count; i++ {
				disk = append(disk, id)
			}

			if n+2 >= len(line) {
				break
			}
			free, err := strconv.Atoi(line[n+1 : n+2])
			if err != nil {
				log.Fatalf("could not parse free %q at %d: %v", line[n+1:n+2], n, err)
			}
			for i := 0; i < free; i++ {
				disk = append(disk, -1)
			}
			id++
		}

		fmt.Println(disk)
		last := len(disk) - 1
		first := 0
		for {
			for ; last >= 0; last-- {
				if disk[last] != -1 {
					break
				}
			}
			for ; first < len(disk); first++ {
				if disk[first] == -1 {
					break
				}
			}
			if first >= len(disk) || last < 0 || first >= last {
				break
			}

			disk[first] = disk[last]
			disk[last] = -1

			// fmt.Println(disk)
		}

		for n, s := range disk {
			if s == -1 {
				break
			}
			sum += n * s
		}
	}

	fmt.Println("Answer:", sum)
}
