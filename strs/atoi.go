package strs

import (
	"log"
	"strconv"
)

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatalf("could not parse int from '%s': %#v", a, err)
	}
	return i
}
