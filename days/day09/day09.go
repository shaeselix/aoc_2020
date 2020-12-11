package day09

import (
	"fmt"
	"aoc_2020/utils"
	"aoc_2020/days/day01"
)

const (
	preamble = 25
)

func firstNonValid(is []int, n int) int {
	for i := range is {
		a, b := day01.FindPair(is[i:i+n], is[i+n])
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

func csProduct(cs []int) int {
	max := cs[0]
	min := cs[0]
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
	csp := csProduct(cs)
	fmt.Println(csp)
}
