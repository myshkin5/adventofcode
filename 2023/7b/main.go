package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

var cardValues = map[rune]int{
	'T': 10,
	'J': 1,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var hands []hand
	for s.Scan() {
		hands = append(hands, parseHand(string(s.Bytes())))
	}

	slices.SortFunc(hands, func(a, b hand) int {
		if a.hType != b.hType {
			return int(a.hType - b.hType)
		}
		for n := 0; n < 5; n++ {
			ac := a.cards[n]
			bc := b.cards[n]
			if ac == bc {
				continue
			}
			aVal, ok := cardValues[rune(ac)]
			if !ok {
				aVal = strs.Atoi(string(ac))
			}
			bVal, ok := cardValues[rune(bc)]
			if !ok {
				bVal = strs.Atoi(string(bc))
			}
			return aVal - bVal
		}
		return 0
	})

	winnings := 0
	for n, hand := range hands {
		fmt.Printf("%s %d %d %d\n", hand.cards, hand.jokers, hand.hType, hand.bid)
		winnings += (n + 1) * hand.bid
	}

	fmt.Println("Answer:", winnings)
}

type hand struct {
	cards  string
	bid    int
	hType  handType
	jokers int
}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func parseHand(line string) hand {
	s1 := strings.Split(line, " ")
	if len(s1) != 2 {
		log.Fatalf("not exactly two sections at ' ': %s", line)
	}
	cards := s1[0]

	hType, jokers := doHandType(cards)
	return hand{
		cards:  cards,
		bid:    strs.Atoi(s1[1]),
		hType:  hType,
		jokers: jokers,
	}
}

func doHandType(cards string) (handType, int) {
	if len(cards) != 5 {
		log.Fatalf("bad cards: %s", cards)
	}
	counts := make(map[rune]int)
	for _, c := range cards {
		counts[c] = counts[c] + 1
	}

	jokers, ok := counts['J']
	if ok {
		delete(counts, 'J')
	}

	if len(counts) <= 1 {
		return fiveOfAKind, jokers
	}
	if len(counts) == 2 {
		for _, v := range counts {
			if v+jokers == 4 {
				return fourOfAKind, jokers
			}
		}
		return fullHouse, jokers
	}
	if len(counts) == 3 {
		for _, v := range counts {
			if v+jokers == 3 {
				return threeOfAKind, jokers
			}
		}
		return twoPair, jokers
	}
	if len(counts) == 4 {
		return onePair, jokers
	}

	if len(counts) != 5 {
		log.Fatalf("wtf: %s", cards)
	}
	return highCard, jokers
}
