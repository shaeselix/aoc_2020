package day3

import (
	"fmt"
	"io/ioutil"
	"aoc_2020/utils"
)

type board [][]uint8

type point struct {
	x int
	y int
}

type slope struct {
	dx int
	dy int
}

var (
	Part2Slopes = []slope {
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}
)

func strToRow(s string) []uint8 {
	row := make([]uint8, len(s))
	for i, v := range s {
		if v == '#' {
			row[i] = uint8(1)
		} else {
			row[i] = uint8(0)
		}
	}
	return row
}

func strArrayToBoard(strs []string) board {
	b := board(make([][]uint8, len(strs)))
	for i, v := range strs {
		b[i] = strToRow(v)
	}
	return b
}

func (b board) traverse(s slope) []point {
	maxX := len(b[0])
	maxY := len(b)
	ps := []point{{0, 0}}
	j := 0
	for i := s.dy;i < maxY;i += s.dy {
		j += s.dx
		if j >= maxX {
			j = j % maxX
		}
		ps = append(ps, point{x: j, y: i})
	}
	return ps
}

func (b board) countTrees(ps []point) int {
	sum := 0
	for _, p := range ps {
		sum += int(b[p.y][p.x])
	}
	return sum
}

func multiSlopeProduct(b board, ss []slope) int {
	p := 1
	for _, s := range ss {
		ps := b.traverse(s)
		ts := b.countTrees(ps)
		p *= ts
	}
	return p
}

func Execute(fp string) {
	dat, err := ioutil.ReadFile(fp)
	utils.Check(err)
	strs := utils.StrArrayFromBytes(dat, "\n")
	b := strArrayToBoard(strs)
	ps := b.traverse(slope{3, 1})
	ts := b.countTrees(ps)
	fmt.Println(ts)
	msp := multiSlopeProduct(b, Part2Slopes)
	fmt.Println(msp)
}