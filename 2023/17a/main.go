package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"

	"github.com/myshkin5/adventofcode/paths"
	"github.com/myshkin5/adventofcode/strs"
)

type loc struct {
	x, y int
}

type direction int

const (
	up direction = 1 << iota
	right
	down
	left
)

func (d direction) String() string {
	switch d {
	case up:
		return "up"
	case right:
		return "right"
	case down:
		return "down"
	case left:
		return "left"
	default:
		return "start"
	}
}

// type possible struct {
// 	loss   int
// 	l      loc
// 	parent *possible
// 	dir    direction
// }

type check struct {
	l   loc
	dir direction
}

func main() {
	f, err := os.Open(filepath.Join(paths.SourcePath(), "input.txt"))
	if err != nil {
		log.Fatalf("could not open file: %#v", err)
	}
	s := bufio.NewScanner(f)

	var grid [][]int
	for s.Scan() {
		line := string(s.Bytes())

		var nums []int
		for _, r := range line {
			nums = append(nums, strs.Atoi(string(r)))
		}

		grid = append(grid, nums)
	}

	// printSteps := func(p *possible) {
	// 	var steps []string
	// 	prevDir := "done"
	// 	for {
	// 		steps = append([]string{fmt.Sprintf("X: %d, y: %d, dir: %s, loss: %d",
	// 			p.l.x+1, p.l.y+1, p.dir, p.loss)}, steps...)
	// 		p = p.parent
	// 		if p == nil {
	// 			break
	// 		}
	// 		prevDir = p.dir.String()
	// 		_ = prevDir
	// 	}
	//
	// 	for _, step := range steps {
	// 		fmt.Println(step)
	// 	}
	// }

	allDirs := [][]check{
		{
			{dir: up, l: loc{x: 0, y: -1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: left, l: loc{x: -1, y: 0}},
		},
		{
			{dir: left, l: loc{x: -1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: up, l: loc{x: 0, y: -1}},
		},
		{
			{dir: up, l: loc{x: 0, y: -1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: left, l: loc{x: -1, y: 0}},
		},
		{
			{dir: left, l: loc{x: -1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: up, l: loc{x: 0, y: -1}},
		},
		{
			{dir: up, l: loc{x: 0, y: -1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: left, l: loc{x: -1, y: 0}},
		},
		{
			{dir: left, l: loc{x: -1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: up, l: loc{x: 0, y: -1}},
		},
		{
			{dir: up, l: loc{x: 0, y: -1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: left, l: loc{x: -1, y: 0}},
		},
		{
			{dir: left, l: loc{x: -1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: up, l: loc{x: 0, y: -1}},
		},
		{
			{dir: up, l: loc{x: 0, y: -1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: left, l: loc{x: -1, y: 0}},
		},
		{
			{dir: left, l: loc{x: -1, y: 0}},
			{dir: down, l: loc{x: 0, y: 1}},
			{dir: right, l: loc{x: 1, y: 0}},
			{dir: up, l: loc{x: 0, y: -1}},
		},
	}
	for n := range allDirs {
		rand.Shuffle(4, func(i, j int) {
			allDirs[n][i], allDirs[n][j] = allDirs[n][j], allDirs[n][i]
		})
	}

	lowestLoss := uint32(1000)
	// var lowestP *possible
	var move func(wId, depth, loss int, l loc, dir1, dir2, dir3 direction, path map[loc]struct{})
	move = func(wId, depth, loss int, l loc, dir1, dir2, dir3 direction, path map[loc]struct{}) {
		if l.x < 0 || l.y < 0 || l.x >= len(grid[0]) || l.y >= len(grid) {
			return
		}
		loss = loss + grid[l.y][l.x]
		ll := atomic.LoadUint32(&lowestLoss)
		if loss > int(ll) {
			return
		}

		if l.x == len(grid[0])-1 && l.y == len(grid)-1 {
			if loss < int(ll) {
				atomic.SwapUint32(&lowestLoss, uint32(loss))
				// lowestP = parent
				log.Println("Solved depth:", depth, "loss:", loss)
			}
			return
		}

		if _, ok := path[l]; ok {
			return
		}

		// cur := &possible{loss: loss, l: l, parent: parent, dir: dir4}
		sameDirs := false
		if dir3 == dir2 && dir2 == dir1 && dir3 != 0 {
			sameDirs = true
		}

		path[l] = struct{}{}

		for _, d := range allDirs[(depth+wId)%len(allDirs)] {
			if !sameDirs || dir3 != d.dir {
				move(wId, depth+1, loss, loc{x: l.x + d.l.x, y: l.y + d.l.y}, dir2, dir3, d.dir, path)
			}
		}

		delete(path, l)
	}

	// printSteps(lowestP)

	wg := sync.WaitGroup{}
	for n := 0; n < 6; n++ {
		wg.Add(1)
		go func(wId int) {
			path := map[loc]struct{}{}
			move(wId, 0, -grid[0][0], loc{x: 0, y: 0}, 0, 0, 0, path)
			wg.Done()
		}(n)
	}

	wg.Wait()

	fmt.Println("Answer:", lowestLoss)
}
