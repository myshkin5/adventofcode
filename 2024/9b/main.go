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

		type span struct {
			start, length int
		}
		files := make(map[int]span)
		var freeSpans []span
		var disk []int
		id := 0
		for n := 0; n < len(line); n += 2 {
			count, err := strconv.Atoi(line[n : n+1])
			if err != nil {
				log.Fatalf("could not parse %q at %d: %v", line[n:n+1], n, err)
			}
			files[id] = span{start: len(disk), length: count}
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
			freeSpans = append(freeSpans, span{start: len(disk), length: free})
			for i := 0; i < free; i++ {
				disk = append(disk, -1)
			}
			id++
			// if id > 10 {
			// 	break
			// }
		}
		maxId := id

		toStr := func() string {
			sb := strings.Builder{}
			prev := -1
			for _, d := range disk {
				if prev != d {
					sb.WriteString("|")
					prev = d
				}
				s := "."
				if d != -1 {
					s = strconv.Itoa(d)
				}
				if _, err := sb.WriteString(s); err != nil {
					log.Fatalf("could not serialize disk %d: %v", d, err)
				}
			}
			return sb.String() + "\n\n\n"
		}

		fmt.Println(toStr())
		for id := maxId; id >= 0; id-- {
			file := files[id]
			fsIndex := -1
			for n, fs := range freeSpans {
				if fs.length >= file.length && fs.start < file.start {
					fsIndex = n
					break
				}
			}
			if fsIndex == -1 {
				continue
			}
			fs := freeSpans[fsIndex]
			for f := fs.start; f < fs.start+file.length; f++ {
				disk[f] = id
			}
			for f := file.start; f < file.start+file.length; f++ {
				disk[f] = -1
			}
			if fs.length == file.length {
				freeSpans = append(freeSpans[:fsIndex], freeSpans[fsIndex+1:]...)
			} else {
				freeSpans[fsIndex] = span{
					start:  fs.start + file.length,
					length: fs.length - file.length,
				}
			}
		}

		fmt.Println(toStr())

		for n, s := range disk {
			if s == -1 {
				continue
			}
			sum += n * s
		}
	}

	fmt.Println("Answer:", sum)
}
