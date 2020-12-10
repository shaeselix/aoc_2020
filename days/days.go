package days

import (
	"aoc_2020/utils"
	"aoc_2020/days/day1"
	"aoc_2020/days/day2"
	"aoc_2020/days/day3"
	"aoc_2020/days/day4"
	"aoc_2020/days/day5"
	"aoc_2020/days/day6"
	"aoc_2020/days/day7"
	"aoc_2020/days/day8"
	"aoc_2020/days/day9"
)

var (
	DayFns = map[int]func(string) {
		1: day1.Execute,
		2: day2.Execute,
		3: day3.Execute,
		4: day4.Execute,
		5: day5.Execute,
		6: day6.Execute,
		7: day7.Execute,
		8: day8.Execute,
		9: day9.Execute,
	}
)

func ExecuteDay(day int, datadir string) {
	utils.DownloadFile(day, datadir, false)
	fp := utils.GetFileName(day, datadir)
	DayFns[day](fp)
}
