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
	cts := map[int]int{0: 1}
	for j, x := range sis {
		i := j - maxdiff
		if i <= 0 {
			i = 0
		}
		s := 0
		for i < j {
			if x - sis[i] <= maxdiff {
				s += cts[i]
			}
			i++
		}
		if j > 0 {
			cts[j] = s
		}
	}
	return cts[len(sis)-1]
}

func Execute(fp string) {
	is := utils.ReadFileAsIntArray(fp, "\n")
	nd := ndiffs(is)
	fmt.Println(nd)
	nc := ncombs(is, 3)
	fmt.Println(nc)
}
