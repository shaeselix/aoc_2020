package day10

import (
	"fmt"
	"aoc_2020/utils"
	"sort"
)

func ndiffs(is []int) int {
	sis := sort.IntSlice(is)
	sis.Sort()
	w := 0
	n1 := 0
	n3 := 1
	for _, x := range sis {
		if x - w == 1 {
			n1++
		}
		if x - w == 3 {
			n3++
		}
		w = x
	}
	return n1 * n3
}

func ncombs(is []int, maxdiff int) int {
	sis := sort.IntSlice(is)
	sis.Sort()
	sis = append([]int{0}, sis...)
	var recur func(i int, j int, cts map[int]int) int
	recur = func(i int, j int, cts map[int]int) int {
		if j == len(sis) {
			return cts[sis[len(sis)-1]]
		}
		s := 0
		for i < j {
			if sis[j] - sis[i] <= maxdiff {
				s += cts[sis[i]]
			}
			i++
		}
		cts[sis[j]] = s
		j++
		ni := j - maxdiff
		if ni <= 0 {
			ni = 0
		}
		return recur(ni, j, cts)
	}
	return recur(0, 1, map[int]int{0: 1})
}

func Execute(fp string) {
	is := utils.ReadFileAsIntArray(fp, "\n")
	nd := ndiffs(is)
	fmt.Println(nd)
	nc := ncombs(is, 3)
	fmt.Println(nc)
}
