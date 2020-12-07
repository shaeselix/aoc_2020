package day7

import (
	"aoc_2020/utils"
	"fmt"
	"strings"
	"strconv"
)

type color string

type bag struct {
	clr		color
	cns		map[color]int
}

type tree struct {
	lvs			map[color]int
	pts			[]color
}

func strToBag(s string) bag {
	cnssplt := strings.Split(s, " bags contain ")
	clr := color(cnssplt[0])
	cns := map[color]int{}
	if cnssplt[1] == "no other bags." {
		return bag{clr, cns}
	}
	commasplit := strings.Split(cnssplt[1], ",")
	for _, v := range commasplit {
		i, c := strToContains(v)
		cns[c] = i
	}
	return bag{clr, cns}
}

func strToContains(s string) (int, color) {
	bgsplit := strings.Split(s, " bag")
	isplit := strings.SplitN(strings.TrimSpace(bgsplit[0]), " ", 2)
	i, err := strconv.Atoi(isplit[0])
	utils.Check(err)
	return i, color(isplit[1])
}

func fileToBags(strs []string) []bag {
	bm := []bag{}
	for _, v := range strs {
		bm = append(bm, strToBag(v))
	}
	return bm
}

func buildTrees(bm []bag) map[color]tree {
	ts := map[color]tree{}
	for _, b := range bm {
		ts[b.clr] = tree{b.cns, []color{}}
	}
	for c, t := range ts {
		ts = fillTree(c, t, ts)
	}
	return ts
}

func fillTree(tc color, t tree, ts map[color]tree) map[color]tree {
	for c := range t.lvs {
		ti := ts[c]
		ti.pts = append(ti.pts, tc)
		ts[c] = ti
	}
	return ts
}

func nOuterBags(c color, ts map[color]tree) int {
	var recur func(c color, acc map[color]bool) map[color]bool
	recur = func(c color, acc map[color]bool) map[color]bool {
		for _, pc := range ts[c].pts {
			acc[pc] = true
			acc = union(acc, recur(pc, map[color]bool{}))
		}
		return acc
	}
	return len(recur(c, map[color]bool{}))
}

func union(a map[color]bool, b map[color]bool) map[color]bool {
	for c := range b {
		a[c] = true
	}
	return a
}

func nInnerBags(c color, ts map[color]tree) int {
	var recur func(c color, acc int) int
	recur = func(c color, acc int) int {
		for lc, i := range ts[c].lvs {
			acc += i
			acc += i*recur(lc, 0)
		}
		return acc
	}
	return recur(c, 0)
}

func Execute(fp string) {
	strs := utils.ReadFileAsStrArray(fp, "\n")
	bm := fileToBags(strs)
	ts := buildTrees(bm)
	nob := nOuterBags(color("shiny gold"), ts)
	fmt.Println(nob)
	nib := nInnerBags(color("shiny gold"), ts)
	fmt.Println(nib)
}
