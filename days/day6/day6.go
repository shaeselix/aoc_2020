package day6

import (
	"fmt"
	"aoc_2020/utils"
	"strings"
)

type yeses map[rune]bool

func emptyYeses() yeses {
	return yeses(map[rune]bool{})
}

func union(ys []yeses) yeses {
	u := emptyYeses()
	for _, y := range ys {
		for k, _ := range y {
			u[k] = true
		}
	}
	return u
}

func intersect(ys []yeses) yeses {
	if len(ys) == 1 {
		return ys[0]
	}
	i := emptyYeses()
	for k, _ := range ys[0] {
		all := true
		for _, y := range ys[1:] {
			_, ok := y[k]
			if !ok {
				all = false
			}
		}
		if all {
			i[k] = true
		}
	}
	return i
}

func strToYeses(s string) yeses {
	y := emptyYeses()
	for _, v := range s {
		y[v] = true
	}
	return y
}

func familyToYeses(fm string) []yeses {
	ms := strings.Fields(fm)
	ys := []yeses{}
	for _, v := range ms {
		ys = append(ys, strToYeses(v))
	}
	return ys
}

func strArrayToUniqueYeses(strs []string) ([]yeses, []yeses) {
	us := []yeses{}
	is := []yeses{}
	for _, v := range strs {
		fys := familyToYeses(v)
		fuy := union(fys)
		fiy := intersect(fys)
		us = append(us, fuy)
		is = append(is, fiy)
	}
	return us, is
}

func countUniqueAnswers(ys []yeses) int {
	sum := 0
	for _, v := range ys {
		sum += len(v)
	}
	return sum
}

func Execute(fp string) {
	strs := utils.ReadFileAsStrArray(fp, "\n\n")
	uys, iys  := strArrayToUniqueYeses(strs)
	ua := countUniqueAnswers(uys)
	fmt.Println(ua)
	ia := countUniqueAnswers(iys)
	fmt.Println(ia)
}
