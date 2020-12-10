package day9

import (
	"fmt"
	"aoc_2020/utils"
	"aoc_2020/days/day1"
)

const (
	preamble = 25
)

func firstNonValid(is []int, n int) int {
	for i := range is {
		a, b := day1.FindPair(is[i:i+n], is[i+n])
		if a == 0 && b == 0 {
			return is[i+n]
		}
	}
	return 0
}

func contiguousSum(is []int, s int) []int {
	cur := is[0]
	i := 0
	j := 1
	for cur != s {
		if cur < s {
			cur += is[j]
			j++
		} else {
			cur -= is[i]
			i++
		}
	}
	return is[i:j]
}

func csProduct(cs []int, s int) int {
	max := 0
	min := s
	for _, x := range cs {
		if x > max {
			max = x
		}
		if x < min {
			min = x
		}
	}
	return min + max
}

func Execute(fp string) {
	is := utils.ReadFileAsIntArray(fp, "\n")
	nv := firstNonValid(is, preamble)
	fmt.Println(nv)
	cs := contiguousSum(is, nv)
	csp := csProduct(cs, nv)
	fmt.Println(csp)
}
