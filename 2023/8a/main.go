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

	steps := 0
	v, ok := vertices["AAA"]
	if !ok {
		log.Fatalf("did not find AAA start")
	}
	done := false
	tries := 10000
	for n := 0; n < tries; n++ {
		log.Println("Try:", n)
		for _, inst := range insts {
			var next string
			switch inst {
			case 'R':
				next = v.r
			case 'L':
				next = v.l
			default:
				log.Fatalf("unknown direction %s", string(inst))
			}

			v, ok = vertices[next]
			if !ok {
				log.Fatalf("did not find vertex: %s", next)
			}
			steps++

			if next == "ZZZ" {
				done = true
				break
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
