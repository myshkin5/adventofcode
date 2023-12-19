package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	if !s.Scan() {
		log.Fatalf("no route line found in input")
	}
	insts := string(s.Bytes())

	vertices := make(map[string]vertex)
	for s.Scan() {
		line := string(s.Bytes())
		if line == "" {
			continue
		}

		n, v := parseVertex(line)
		vertices[n] = v
	}

	var starts []vertex
	var vs []vertex
	for n, v := range vertices {
		if len(n) < 3 || n[2] != 'A' {
			continue
		}

		starts = append(starts, v)
		vs = append(vs, v)
	}

	steps := 0
	tries := 100_000_000_000
	var done bool
	for n := 0; n < tries; n++ {
		if n%10000000 == 0 {
			log.Println("Try:", n, "steps:", steps)
		}
		for _, inst := range insts {
			done = true
			for m := 0; m < len(vs); m++ {
				v := vs[m]

				var next string
				switch inst {
				case 'R':
					next = v.r
				case 'L':
					next = v.l
				default:
					log.Fatalf("unknown direction %s", string(inst))
				}

				var ok bool
				vs[m], ok = vertices[next]
				if !ok {
					log.Fatalf("did not find vertex: %s", next)
				}

				if len(next) < 3 || next[2] != 'Z' {
					done = false
				}
			}
			steps++

			if done {
				break
			}
		}
		for m := 0; m < len(vs); m++ {
			if vs[m] == starts[m] {
				log.Println("Looped", m, "steps", steps)
			}
		}

		if done {
			break
		}
	}

	if !done {
		log.Fatalf("did not a solution after %d tries", tries)
	}

	fmt.Println("Answer:", steps)
}

type vertex struct {
	l, r string
}

func parseVertex(line string) (string, vertex) {
	s1 := strings.Split(line, " = ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ' = ': %s", line)
	}
	name := s1[0]

	ps := s1[1]
	if ps[0] != '(' || ps[len(ps)-1] != ')' {
		log.Fatalf("no parens found around vertex: %s", line)
	}

	s2 := strings.Split(ps[1:len(ps)-1], ", ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ', ': %s", line)
	}

	return name, vertex{l: s2[0], r: s2[1]}
}
