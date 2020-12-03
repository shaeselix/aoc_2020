package day1

import (
	"fmt"
	"io/ioutil"
	"aoc_2020/utils"
)

const (
	TargetSum int = 2020
)

func findPair(is []int, target int) (int, int) {
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

func removeItem(is []int, item int) []int {
	for i, v := range is {
		if v == item {
			is[i] = is[len(is) - 1]
			return is[:len(is) - 1]
		}
	}
	return is
}

func findTriplet(is []int, target int) (int, int, int) {
	for _, v := range is {
		isv := utils.RemoveItem(is, v)
		o1, o2 := findPair(isv, target - v)
		if o1 != 0 || o2 != 0 {
			return v, o1, o2
		}
	}
	return 0, 0, 0
}

func Execute(fp string) {
	dat, err := ioutil.ReadFile(fp)
	utils.Check(err)
	is := utils.IntArrayFromBytes(dat, "\n")
	x, y := findPair(is, TargetSum)
	fmt.Println(x*y)
	a, b, c := findTriplet(is, TargetSum)
	fmt.Println(a*b*c)
}
