package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

var (
	idRexp = regexp.MustCompile("Game ([0-9]+): (.+)")
	ccRexp = regexp.MustCompile(" *([0-9]+) (red|green|blue)")
	limits = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)
	var sum int
	for s.Scan() {
		line := string(s.Bytes())
		subs := idRexp.FindStringSubmatch(line)
		if len(subs) != 3 {
			log.Fatalf("could not find game id for line: %s", line)
		}
		failed := false
		for _, set := range strings.Split(subs[2], ";") {
			for _, cc := range strings.Split(set, ", ") {
				ccs := ccRexp.FindStringSubmatch(cc)
				if len(ccs) != 3 {
					log.Fatalf("could not parse cube count/color: %s", cc)
				}
				if strs.Atoi(ccs[1]) > limits[ccs[2]] {
					failed = true
				}
			}
		}
		if !failed {
			sum += strs.Atoi(subs[1])
		}
	}

	fmt.Println("Answer:", sum)
}
