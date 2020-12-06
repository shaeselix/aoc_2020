package day5

import (
	"fmt"
	"aoc_2020/utils"
)

type ticket struct {
	row uint8
	col uint8
	sid int
}

const (
	maxRow uint8 = 128
	maxCol uint8 = 8
)

func findTicket(bsp string) ticket {
	row := binSearch(maxRow / 2 - 1, maxRow / 2, bsp[:7], 'B', 'F')
	col := binSearch(maxCol / 2 - 1, maxCol / 2, bsp[7:], 'R', 'L')
	sid := seatId(row, col)
	return ticket{row, col, sid}
}

func binSearch(guess uint8, inc uint8, bsp string, pl byte, nl byte) uint8 {
	if len(bsp) == 1 {
		if bsp[0] == pl {
			return guess + 1
		}
		return guess
	} else {
		inc = inc / 2
		switch step := bsp[0]; step {
		case pl:
			return binSearch(guess + inc, inc, bsp[1:], pl, nl)
		case nl:
			return binSearch(guess - inc, inc, bsp[1:], pl, nl)
		default:
			panic(fmt.Sprintf("unrecognized character: %v", step))
		}
	}
}

func seatId(row uint8, col uint8) int {
	return int(row)*8 + int(col)
}

func getTickets(bsps []string) []ticket {
	ts := []ticket{}
	for _, v := range bsps {
		ts = append(ts, findTicket(v))
	}
	return ts
}

func findHighestSeat(ts []ticket) int {
	max := 0
	for _, v := range ts {
		if v.sid > max {
			max = v.sid
		}
	}
	return max
}

func ticketSliceToSidSet(ts []ticket) map[int]bool {
	ss := map[int]bool{}
	for _, v := range ts {
		ss[v.sid] = true
	}
	return ss
}

func findMissingSeat(ts []ticket, max int) int {
	ss := ticketSliceToSidSet(ts)
	for sid, _ := range ss {
		_, ok := ss[sid + 1]
		if !ok && sid < max {
			return sid + 1
		}
	}
	return 0
}

func Execute(fp string) {
	strs := utils.ReadFileAsStrArray(fp, "\n")
	ts := getTickets(strs)
	hs := findHighestSeat(ts)
	fmt.Println(hs)
	ms := findMissingSeat(ts, hs)
	fmt.Println(ms)
}
