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

	var line string
	for s.Scan() {
		if line != "" {
			log.Fatalf("more than one line")
		}
		line = string(s.Bytes())
	}

	boxes := [256][]slot{}
	for _, step := range strings.Split(line, ",") {
		var subtract bool
		var label string
		var flen int
		if strings.HasSuffix(step, "-") {
			label = step[:len(step)-1]
			subtract = true
		} else {
			s1 := strings.Split(step, "=")
			label = s1[0]
			flen = strs.Atoi(s1[1])
		}
		h := hash(label)
		if subtract {
			slots := boxes[h]
			for n, s := range slots {
				if s.label == label {
					boxes[h] = slots[0:n]
					if n < len(slots)-1 {
						boxes[h] = append(boxes[h], slots[n+1:]...)
					}
					break
				}
			}
		} else {
			updated := false
			for n, s := range boxes[h] {
				if s.label == label {
					boxes[h][n].flen = flen
					updated = true
					break
				}
			}
			if !updated {
				boxes[h] = append(boxes[h], slot{label: label, flen: flen})
			}
		}
	}

	total := 0
	for n, box := range boxes {
		for m, s := range box {
			add := (n + 1) * (m + 1) * s.flen
			fmt.Println("Box:", n, box, "adds:", add)
			total += add
		}
	}

	fmt.Println("Answer:", total)
}

type slot struct {
	label string
	flen  int
}

func hash(step string) int {
	val := 0
	for _, r := range step {
		val += int(r)
		val *= 17
		val %= 256
	}
	return val
}
