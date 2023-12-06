package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	if !s.Scan() {
		log.Fatalf("no lines found in input")
	}
	seedRanges := parseSeeds(string(s.Bytes()))
	ranges := parseRanges(s)

	minLoc := math.MaxInt
	for _, seedRange := range seedRanges {
		log.Printf("Processing seed range %d with length %d", seedRange.start, seedRange.len)
		for n := seedRange.start; n < seedRange.start+seedRange.len; n++ {
			loc := findLocation(n, ranges)
			if loc < minLoc {
				minLoc = loc
			}
		}
	}

	fmt.Println("Answer:", minLoc)
}

type seedRange struct {
	start, len int
}

func parseSeeds(line string) []seedRange {
	s1 := strings.Split(line, ": ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ': ': %s", line)
	}
	if s1[0] != "seeds" {
		log.Fatalf("didn't find seeds label: %s", line)
	}

	s2 := strings.Split(s1[1], " ")
	if len(s2)%2 != 0 {
		log.Fatalf("seed ranges not given as pairs: %s", line)
	}
	var seeds []seedRange
	for n := 0; n < len(s2)-1; n += 2 {
		seeds = append(seeds, seedRange{start: strs.Atoi(s2[n]), len: strs.Atoi(s2[n+1])})
	}
	return seeds
}

type mapRange struct {
	dest, source, len int
}

func parseRanges(s *bufio.Scanner) [][]mapRange {
	var ranges [][]mapRange
	var from, to string
	for s.Scan() {
		label := s.Bytes()
		if len(label) == 0 {
			continue
		}

		s1 := strings.Split(string(label), " ")
		if len(s1) != 2 {
			log.Fatalf("could not parse label at ' ': %s", label)
		}
		if s1[1] != "map:" {
			log.Fatalf("unexpected label format: %s", label)
		}

		s2 := strings.Split(s1[0], "-")
		if len(s2) != 3 {
			log.Fatalf("could not parse from/to: %s", label)
		}
		if s2[1] != "to" {
			log.Fatalf("unexpected from/to format: %s", label)
		}
		if from == "" || to == "" {
			if s2[0] != "seed" {
				log.Fatalf("first map must be from seed, found %s", s2[0])
			}
		} else {
			if to != s2[0] {
				log.Fatalf("maps not in order from %s-to-%s to %s-to-%s", from, to, s2[0], s2[2])
			}
		}
		from = s2[0]
		to = s2[2]

		var subRanges []mapRange
		for s.Scan() {
			line := string(s.Bytes())
			if line == "" {
				break
			}

			s3 := strings.Split(line, " ")
			if len(s3) != 3 {
				log.Fatalf("could not read range in %s-to-%s: %s", from, to, line)
			}

			subRanges = append(subRanges, mapRange{
				dest:   strs.Atoi(s3[0]),
				source: strs.Atoi(s3[1]),
				len:    strs.Atoi(s3[2]),
			})
		}

		log.Printf("Parsed %s-to-%s with %d ranges\n", from, to, len(ranges))
		ranges = append(ranges, subRanges)
	}

	if to != "location" {
		log.Fatalf("last map must be to location, found %s", to)
	}

	return ranges
}

func findLocation(id int, ranges [][]mapRange) int {
	for _, subRanges := range ranges {
		for _, subRange := range subRanges {
			if id >= subRange.source && id < subRange.source+subRange.len {
				id = subRange.dest + id - subRange.source
				break
			}
		}
		// if !found, id is used in next map
	}
	return id
}
