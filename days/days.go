package days

import (
	"aoc_2020/utils"
	"aoc_2020/days/day01"
	"aoc_2020/days/day02"
	"aoc_2020/days/day03"
	"aoc_2020/days/day04"
	"aoc_2020/days/day05"
	"aoc_2020/days/day06"
	"aoc_2020/days/day07"
	"aoc_2020/days/day08"
	"aoc_2020/days/day09"
	"aoc_2020/days/day10"
	"aoc_2020/days/day11"
	"aoc_2020/days/day12"
	"aoc_2020/days/day13"
)

var (
	DayFns = map[int]func(string) {
		1: day01.Execute,
		2: day02.Execute,
		3: day03.Execute,
		4: day04.Execute,
		5: day05.Execute,
		6: day06.Execute,
		7: day07.Execute,
		8: day08.Execute,
		9: day09.Execute,
		10: day10.Execute,
		11: day11.Execute,
		12: day12.Execute,
		13: day13.Execute,
	}
)

func ExecuteDay(day int, datadir string) {
	utils.DownloadFile(day, datadir, false)
	fp := utils.GetFileName(day, datadir)
	DayFns[day](fp)
}
