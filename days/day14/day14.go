package day14

import (
	"fmt"
	"aoc_2020/utils"
	"strings"
	"strconv"
	"math"
)

type mask struct {
	and int
	or int
}

type assign struct {
	index  int
	value  int
}

var test = []string{
	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	"mem[8] = 11",
	"mem[7] = 101",
	"mem[8] = 0",
}

var test2 = []string{
	"mask = 000000000000000000000000000000X1001X",
	"mem[42] = 100",
	"mask = 00000000000000000000000000000000X0XX",
	"mem[26] = 1",
}

func strToMask(s string) mask {
	ands := strings.ReplaceAll(s[7:], "X", "1")
	ors := strings.ReplaceAll(s[7:], "X", "0")
	and, err := strconv.ParseInt(ands, 2, 64)
	utils.Check(err)
	or, err := strconv.ParseInt(ors, 2, 64)
	utils.Check(err)
	return mask{int(and), int(or)}
}

func pow2(power int) int {
	return int(math.Pow(float64(2), float64(power)))
}

func handleX(ors []int, ands []int, power int) ([]int, []int) {
	nors := make([]int, len(ors)*2)
	nands := make([]int, len(ands)*2)
	for i, o := range ors {
		nors[2*i] = o // set to 0
		nors[2*i+1] = o + pow2(power) // set to 1
	}
	for i, a := range ands {
		nands[2*i] = a - pow2(power) // set to 0
		nands[2*i+1] = a // set to 1
	}
	return nors, nands
}

func strToMasks(s string) []mask {
	power := len(s)
	ands := []int{pow2(power) - 1}
	ors := []int{0}
	for _, c := range s {
		power--
		switch c {
		case '1':
			for i, o := range ors {
				ors[i] = o + pow2(power)
			}
		case 'X':
			ors, ands = handleX(ors, ands, power)
		}
	}
	ms := make([]mask, len(ors))
	for i, o := range ors {
		ms[i] = mask{ands[i], o}
	}
	return ms
}

func rowToAssign(s string) assign {
	eqsplit := strings.Split(s, " = ")
	value := utils.StrToInt(eqsplit[1])
	bracsplit := strings.Split(eqsplit[0], "[")
	indstr := bracsplit[1][:len(bracsplit[1]) - 1]
	index := utils.StrToInt(indstr)
	return assign{index, value}
}

func fileApplyMask(ss []string) int {
	mem := map[int]int{}
	m := strToMask(ss[0])
	for _, s := range ss[1:] {
		if s[0:4] == "mask" {
			m = strToMask(s)
		}
		if s[0:3] == "mem" {
			a := rowToAssign(s)
			mem[a.index] = applyMask(a.value, m)
		}
	}
	return sumValues(mem)
}

func fileApplyMasks(ss []string) int {
	mem := map[int]int{}
	ms := strToMasks(ss[0])
	for _, s := range ss[1:] {
		if s[0:4] == "mask" {
			ms = strToMasks(s)
		}
		if s[0:3] == "mem" {
			a := rowToAssign(s)
			for _, m := range ms {
				mem[applyMask(a.index, m)] = a.value
			}
		}
	}
	return sumValues(mem)
}

func applyMask(i int, m mask) int {
	return (i & m.and) | m.or
}

func sumValues(mem map[int]int) int {
	s := 0
	for _, i := range mem {
		s += i
	}
	return s
} 

func Execute(fp string) {
	ss := utils.ReadFileAsStrArray(fp, "\n")
	am := fileApplyMask(ss)
	fmt.Println(am)
	ams := fileApplyMasks(ss)
	fmt.Println(ams)
}
