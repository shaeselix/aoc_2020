package day11

import (
	"fmt"
	"aoc_2020/utils"
	"sync"
)

type cell uint8

const (
	floor cell = iota
	empty
	filled
)

type board [][]cell

func toRow(s string) []cell {
	r := make([]cell, len(s))
	for i, c := range s {
		switch c {
		case 'L':
			r[i] = empty
		case '.':
			r[i] = floor
		}
	}
	return r
}

func fileToBoard(ss []string) board {
	b := board(make([][]cell, len(ss)))
	for i, s := range ss {
		b[i] = toRow(s)
	}
	return b
}

func boundedDiff(a int, b int, min int) int {
	d := a - b
	if d < min {
		d = min
	}
	return d
}

func boundedSum(a int, b int, max int) int {
	s := a + b
	if s > max {
		s = max
	}
	return s
}

func flipcell(r int, c int, mr int, mc int, ob *board) cell {
	oc := (*ob)[r][c]
	if oc == floor {
		return oc
	}
	mi := boundedSum(r, 2, mr)
	mj := boundedSum(c, 2, mc)
	s := 0
	i := boundedDiff(r, 1, 0)
	for i < mi {
		j := boundedDiff(c, 1, 0)
		for j < mj {
			if (i != r || j != c) && (*ob)[i][j] == filled {
				s++
				if s >= 4 {
					return empty
				}
			}
			j ++ 
		}
		i++
	}
	if s == 0 {
		return filled
	}
	return oc
}

func flipcellp2(r int, c int, mr int, mc int, ob *board) cell {
	oc := (*ob)[r][c]
	if oc == floor {
		return oc
	}
	s := 0
	for i := -1;i < 2;i++ {
		for j := -1;j < 2;j++ {
			if (i != 0 || j != 0) {
				fs := findseat(r, c, mr, mc, ob, i, j)
				if fs == filled {
					s++
					if s >= 5 {
						return empty
					}
				}
			}
		}
	}
	if s == 0 {
		return filled
	}
	return oc
}

func findseat(r int, c int, mr int, mc int, ob *board, dy int, dx int) cell {
	for {
		r += dy
		c += dx
		if r < 0 || r >= mr || c < 0 || c >= mc {
			return floor
		}
		if (*ob)[r][c] != floor {
			return (*ob)[r][c]
		}
	}
}

func parStep(b board, fc func(int, int, int, int, *board) cell) board {
	y := len(b)
	x := len(b[0])
	nb := board(make([][]cell, y))
	for i := range nb {
		nb[i] = make([]cell, x)
	}
	var wg sync.WaitGroup
	wg.Add(x*y)
	for r := range b {
		for c := range b[r] {
			go func(r int, c int) {
				defer wg.Done()
				nb[r][c] = fc(r, c, y, x, &b)
			}(r, c)
		}
	}
	wg.Wait()
	return nb
}

func (b board) equals(ob board) bool {
	for i, r := range b {
		for j, x := range r {
			if x != ob[i][j] {
				return false
			}
		}
	}
	return true
}

func noccupied(b board) int {
	cnt := 0
	for _, r := range b {
		for _, x := range r {
			if x == filled {
				cnt++
			}
		}
	}
	return cnt
}

func stabilize(b board, maxi int, fc func(int, int, int, int, *board) cell) int {
	for i := 0;i < maxi;i++ {
		nb := parStep(b, fc)
		if b.equals(nb) {
			return noccupied(b)
		}
		b = nb
	}
	return 0
}

func Execute(fp string) {
	ss := utils.ReadFileAsStrArray(fp, "\n")
	b := fileToBoard(ss)
	sb := stabilize(b, 1000000000, flipcell)
	fmt.Println(sb)
	sb2 := stabilize(b, 1000000000, flipcellp2)
	fmt.Println(sb2)
}
