package day08

import (
	"aoc_2020/utils"
	"fmt"
	"strings"
	"strconv"
)

type instruction struct {
	op  string
	arg int
}

func toInstruction(s string) instruction {
	sp := strings.Fields(s)
	var stri string;
	if sp[1][0] == '+' {
		stri = sp[1][1:]
	} else {
		stri = sp[1]
	}
	i, err := strconv.Atoi(stri)
	utils.Check(err)
	return instruction{sp[0], i}
}

func fileToInstructions(strs []string) []instruction {
	is := []instruction{}
	for _, v := range strs {
		is = append(is, toInstruction(v))
	}
	return is
}

func step(inst instruction, i int, acc int) (int, int) {
	switch inst.op {
	case "acc":
		return i + 1, acc + inst.arg
	case "jmp":
		return i + inst.arg, acc
	case "nop":
		return i + 1, acc
	default:
		fmt.Println(inst.op)
		return 0, 0
	}
}

func accumulate(is []instruction) (int, int) {
	acc := 0
	previ := 0
	i := 0
	maxi := len(is)
	steps := map[int]bool{}
	for true {
		_, ok := steps[i]
		if ok || i == maxi {
			break;
		} else {
			steps[i] = true
		}
		previ = i
		i, acc = step(is[i], i, acc)
	}
	return acc, previ
}

func switchstep(is []instruction) int {
	nis := make([]instruction, len(is))
	for i, inst := range is {
		if inst.op == "jmp" {
			copy(nis, is)
			nis[i] = instruction{"nop", 0}
			acc, previ := accumulate(nis)
			if previ == len(is) - 1 {
				return acc
			}
		}
	}
	return 0
}

func Execute(fp string) {
	strs := utils.ReadFileAsStrArray(fp, "\n")
	is := fileToInstructions(strs)
	acc, _ := accumulate(is)
	fmt.Println(acc)
	acc = switchstep(is)
	fmt.Println(acc)
}
