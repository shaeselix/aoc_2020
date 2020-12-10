package day1

import (
	"fmt"
	"aoc_2020/utils"
)

const (
	TargetSum int = 2020
)

func FindPair(is []int, target int) (int, int) {
	sums := make(map[int]int)
	for _, v := range is {
		o, ok := sums[v]
		if ok {
			return v, o
		} else {
			sums[target - v] = v
		}
	}
	return 0, 0
}

func findTriplet(is []int, target int) (int, int, int) {
	for _, v := range is {
		isv := utils.RemoveItem(is, v)
		o1, o2 := FindPair(isv, target - v)
		if o1 != 0 || o2 != 0 {
			return v, o1, o2
		}
	}
	return 0, 0, 0
}

func Execute(fp string) {
	is := utils.ReadFileAsIntArray(fp, "\n")
	x, y := FindPair(is, TargetSum)
	fmt.Println(x*y)
	a, b, c := findTriplet(is, TargetSum)
	fmt.Println(a*b*c)
}
