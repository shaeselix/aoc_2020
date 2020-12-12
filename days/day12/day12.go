package day12

import (
	"fmt"
	"aoc_2020/utils"
	"strconv"
)

type action uint8

const (
	N action = iota
	S
	E
	W
	L
	R
	F
)

type instruction struct {
	a action
	v int
}

type position struct {
	ns int
	ew int
	a action
}

func strToInstruction(s string) instruction {
	var a action
	switch s[0] {
	case 'N':
		a = N
	case 'S':
		a = S
	case 'E':
		a = E
	case 'W':
		a = W
	case 'L':
		a = L
	case 'R':
		a = R
	case 'F':
		a = F
	}
	v, err := strconv.Atoi(s[1:])
	utils.Check(err)
	return instruction{a, v}
}

func fileToInstructions(ss []string) []instruction {
	is := make([]instruction, len(ss))
	for i, s := range ss {
		is[i] = strToInstruction(s)
	}
	return is
}

func north(p position, v int) position {
	return position{p.ns + v, p.ew, p.a}
}

func south(p position, v int) position {
	return position{p.ns - v, p.ew, p.a}
}

func east(p position, v int) position {
	return position{p.ns, p.ew + v, p.a}
}

func west(p position, v int) position {
	return position{p.ns, p.ew - v, p.a}
}

func left(p position, v int) position {
	for i := 0;i<v/90;i++ {
		switch p.a {
		case N:
			p.a = W
		case S:
			p.a = E
		case E:
			p.a = N
		case W:
			p.a = S
		}
	}
	return p
}

func right(p position, v int) position {
	for i := 0;i<v/90;i++ {
		switch p.a {
		case N:
			p.a = E
		case S:
			p.a = W
		case E:
			p.a = S
		case W:
			p.a = N
		}
	}
	return p
}

func forward(p position, v int) position {
	switch p.a {
	case N:
		return position{p.ns + v, p.ew, N}
	case S:
		return position{p.ns - v, p.ew, S}
	case E:
		return position{p.ns, p.ew + v, E}
	case W:
		return position{p.ns, p.ew - v, W}
	default:
		return p
	}
}

func step(p position, i instruction) position {
	switch i.a {
	case N:
		return north(p, i.v)
	case S:
		return south(p, i.v)
	case E:
		return east(p, i.v)
	case W:
		return west(p, i.v)
	case L:
		return left(p, i.v)
	case R:
		return right(p, i.v)
	case F:
		return forward(p, i.v)
	default:
		return p
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func manhattan(p position) int {
	return abs(p.ns) + abs(p.ew)
}

func ship(p position, is []instruction) position {
	for _, i := range is {
		p = step(p, i)
	}
	return p
}

func wayforward(p position, wp position, v int) position {
	return position{p.ns + wp.ns*v, p.ew + wp.ew*v, p.a}
}

func wayleft(wp position, v int) position {
	for i := 0;i < v/90;i++ {
		wp = position{wp.ew, -wp.ns, wp.a}
	}
	return wp
}

func wayright(wp position, v int) position {
	for i := 0;i < v/90;i++ {
		wp = position{-wp.ew, wp.ns, wp.a}
	}
	return wp
}

func waystep(p position, wp position, i instruction) (position, position) {
	switch i.a {
	case N:
		return p, north(wp, i.v)
	case S:
		return p, south(wp, i.v)
	case E:
		return p, east(wp, i.v)
	case W:
		return p, west(wp, i.v)
	case L:
		return p, wayleft(wp, i.v)
	case R:
		return p, wayright(wp, i.v)
	case F:
		return wayforward(p, wp, i.v), wp
	default:
		return p, wp
	}
}

func wayship(p position, wp position, is []instruction) position {
	for _, i := range is {
		p, wp = waystep(p, wp, i)
	}
	return p
}

func Execute(fp string) {
	ss := utils.ReadFileAsStrArray(fp, "\n")
	is := fileToInstructions(ss)
	p0 := position{0, 0, E}
	pf := ship(p0, is)
	fmt.Println(manhattan(pf))
	wp0 := position{1, 10, N}
	pfw := wayship(p0, wp0, is)
	fmt.Println(manhattan(pfw))
}
