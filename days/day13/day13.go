package day13

import (
	"fmt"
	"aoc_2020/utils"
	"strings"
	"math"
)

const (
	test = "7,13,x,x,59,x,31,19"
)

type bus struct {
	bid int
	ind int
}

func strToBuses(s string) []bus {
	sps := strings.Split(s, ",")
	bs := []bus{}
	for j, sp := range sps {
		if sp != "x" {
			b := bus{utils.StrToInt(sp), j}
			bs = append(bs, b)
		}
	}
	return bs
}

func wait(t int, b int) int {
	if t % b == 0 {
		return 0
	}
	return b - t % b
}

func soonestBus(t int, bs []bus) (int, int) {
	mw := math.MaxInt32
	bb := 0
	for _, b := range bs {
		w := wait(t, b.bid)
		if w <= mw {
			mw = w
			bb = b.bid
		}
	}
	return bb, mw
}

func check(b bus, inc int, x int) int {
	for {
		if (x + b.ind) % b.bid == 0 {
			return x
		}
		x += inc
	}
}

func itercon(bs []bus) int {
	inc := 1
	n := 1
	for _, b := range bs {
		n = check(b, inc, n)
		inc *= b.bid
	}
	return n
}

func Execute(fp string) {
	ss := utils.ReadFileAsStrArray(fp, "\n")
	t := utils.StrToInt(ss[0])
	bs := strToBuses(ss[1])
	bb, mw := soonestBus(t, bs)
	fmt.Println(bb*mw)
	c := itercon(bs)
	fmt.Println(c)
}
