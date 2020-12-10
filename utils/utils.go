package utils

import (
	"strings"
	"strconv"
	"io/ioutil"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func StrArrayFromBytes(dat []byte, splitstr string) []string {
	stripped := strings.TrimSpace(string(dat))
	return strings.Split(stripped, splitstr)
}

func IntArrayFromBytes(dat []byte, splitstr string) []int {
	strs := StrArrayFromBytes(dat, splitstr)
	is := make([]int, 0)
	for _, v := range strs {
		if len(v) > 0 {
			i, err := strconv.Atoi(v)
			Check(err)
			is = append(is, i)
		}
	}
	return is
}

func ReadFileAsStrArray(fp string, splitstr string) []string {
	dat, err := ioutil.ReadFile(fp)
	Check(err)
	return StrArrayFromBytes(dat, splitstr)
}

func ReadFileAsIntArray(fp string, splitstr string) []int {
	dat, err := ioutil.ReadFile(fp)
	Check(err)
	return IntArrayFromBytes(dat, splitstr)
}

func RemoveItem(is []int, item int) []int {
	for i, v := range is {
		if v == item {
			is[i] = is[len(is) - 1]
			return is[:len(is) - 1]
		}
	}
	return is
}
