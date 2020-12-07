package days

import (
	"fmt"
	"aoc_2020/days/day1"
	"aoc_2020/days/day2"
	"aoc_2020/days/day3"
	"aoc_2020/days/day4"
	"aoc_2020/days/day5"
	"aoc_2020/days/day6"
	"aoc_2020/days/day7"
)

var (
	dataFileFormat string = "%d.txt"
	DayFns = map[int]func(string) {
		1: day1.Execute,
		2: day2.Execute,
		3: day3.Execute,
		4: day4.Execute,
		5: day5.Execute,
		6: day6.Execute,
		7: day7.Execute,
	}
)

func ExecuteDay(day int, datadir string) {
	file := fmt.Sprintf(dataFileFormat, day)
	fp := fmt.Sprintf("%s%s", datadir, file)
	DayFns[day](fp)
}
