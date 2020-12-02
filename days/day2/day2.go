package day2

import (
	"fmt"
	"io/ioutil"
	"aoc_2020/utils"
	"strings"
	"strconv"
)

var (
	FilePath string = "days/day2/data/input.txt"
)

type Policy struct {
	password string
	letter string
	min int
	max int
}

func strToPolicy(s string) Policy {
	dashSplit := strings.Split(s, "-")
	minStr := dashSplit[0]
	colonSplit := strings.Split(dashSplit[1], ": ")
	pwStr := colonSplit[1]
	spaceSplit := strings.Split(colonSplit[0], " ")
	maxStr := spaceSplit[0]
	letterStr := spaceSplit[1]
	min, err := strconv.Atoi(minStr)
	utils.Check(err)
	max, err := strconv.Atoi(maxStr)
	utils.Check(err)
	return Policy{
		password: pwStr,
		letter: letterStr,
		min: min, 
		max: max}
}

func strArrayToPolicies(strs []string) []Policy {
	ps := make([]Policy, len(strs))
	for i, v := range strs {
		ps[i] = strToPolicy(v)
	}
	return ps
}

func isValidP1(p Policy) bool {
	count := strings.Count(p.password, p.letter)
	if count < p.min {
		return false
	}
	if count > p.max {
		return false
	}
	return true
}

func countValids(ps []Policy, pf func(Policy) bool) int {
	sum := 0
	for _, v := range ps {
		if pf(v) {
			sum ++
		}
	}
	return sum
}

func isValidP2(p Policy) (b bool) {
	defer func() {
		if r := recover(); r!= nil {
			b = false
		}
	}()
	p1 := p.password[p.min-1] == p.letter[0]
	p2 := p.password[p.max-1] == p.letter[0]
	return p1 != p2
}

func Execute() {
	dat, err := ioutil.ReadFile(FilePath)
	utils.Check(err)
	strs := utils.StrArrayFromBytes(dat, "\n")
	ps := strArrayToPolicies(strs)
	fmt.Println(countValids(ps, isValidP1))
	fmt.Println(countValids(ps, isValidP2))
}
