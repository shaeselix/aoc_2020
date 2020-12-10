package day10

import (
	"fmt"
	"aoc_2020/utils"
)

func ndiffs(sis []int) int {
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

func ncombs(sis []int, maxdiff int) int {
	sis = append([]int{0}, sis...)
	cts := []int{1}
	for j := 1;j < len(sis);j++ {
		cts = append(cts, combstep(sis, j, maxdiff, cts))
	}
	return cts[len(sis)-1]
}

func combstep(sis []int, j int, maxdiff int, cts []int) int {
	i := j - maxdiff
	if i <= 0 {
		i = 0
	}
	s := 0
	for i < j {
		if sis[j] - sis[i] <= maxdiff {
			s += cts[i]
		}
		i++
	}
	return s
}

func Execute(fp string) {
	is := utils.ReadFileAsIntArray(fp, "\n")
	sis := utils.SortIntArray(is)
	nd := ndiffs(sis)
	fmt.Println(nd)
	nc := ncombs(sis, 3)
	fmt.Println(nc)
}
