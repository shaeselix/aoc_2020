package day4

import (
	"fmt"
	"io/ioutil"
	"aoc_2020/utils"
	"strings"
	"strconv"
	"regexp"
)

var (
	requiredFields = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}
)

type passport map[string]string

type rfpassport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
}

func strToPassport(s string) passport {
	p := passport(make(map[string]string))
	for _, v := range strings.Fields(s) {
		sp := strings.Split(v, ":")
		p[sp[0]] = sp[1]
	}
	return p
}

func strArrayToPassports(strs []string) []passport {
	ps := []passport{}
	for _, v := range strs {
		ps = append(ps, strToPassport(v))
	}
	return ps
}

func (p passport) hasRequiredFields() bool {
	for _, f := range requiredFields {
		_, ok := p[f]
		if !ok {
			return false
		}
	}
	return true
}

func (p passport) toRFP() rfpassport {
	return rfpassport{
		byr: p["byr"],
		iyr: p["iyr"],
		eyr: p["eyr"],
		hgt: p["hgt"],
		hcl: p["hcl"],
		ecl: p["ecl"],
		pid: p["pid"],
	}
}

func countPassportsWithRequiredFields(ps []passport) (int, []rfpassport) {
	count := 0
	rfps := []rfpassport{}
	for _, p := range ps {
		if p.hasRequiredFields() {
			count += 1
			rfps = append(rfps, p.toRFP())
		}
	}
	return count, rfps
}

func strIntIsValid(si string, min int, max int) bool {
	i, err := strconv.Atoi(si)
	if err != nil { return false }
	if i < min { return false }
	if i > max { return false }
	return true
}

func byrIsValid(byr string) bool {
	return strIntIsValid(byr, 1920, 2002)
}

func iyrIsValid(iyr string) bool {
	return strIntIsValid(iyr, 2010, 2020)
}

func eyrIsValid(eyr string) bool {
	return strIntIsValid(eyr, 2020, 2030)
}

func hgtIsValid(hgt string) bool {
	if strings.Contains(hgt, "cm") {
		si := strings.TrimSuffix(hgt, "cm")
		return strIntIsValid(si, 150, 193)
	}
	if strings.Contains(hgt, "in") {
		si := strings.TrimSuffix(hgt, "in")
		return strIntIsValid(si, 59, 76)
	}
	return false
}

func hclIsValid(hcl string) bool {
	if hcl[0] != '#' { return false }
	if len(hcl) != 7 { return false }
	m, err := regexp.MatchString("[^0-9a-f]", hcl[1:])
	utils.Check(err)
	if m { return false }
	return true
}

func eclIsValid(ecl string) bool {
	ve := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	_, ok := ve[ecl]
	return ok
}

func pidIsValid(pid string) bool {
	if len(pid) != 9 { return false }
	m, err := regexp.MatchString("[^0-9]", pid)
	utils.Check(err)
	if m { return false }
	return true
}

func (p rfpassport) isValid() bool {
	return (
		byrIsValid(p.byr) &&
		iyrIsValid(p.iyr) &&
		eyrIsValid(p.eyr) &&
		hgtIsValid(p.hgt) &&
		hclIsValid(p.hcl) &&
		eclIsValid(p.ecl) &&
		pidIsValid(p.pid) )
}

func countValidPassports(rfps []rfpassport) int {
	count := 0
	for _, p := range rfps {
		if p.isValid() {
			count++
		}
	}
	return count
}

func Execute(fp string) {
	dat, err := ioutil.ReadFile(fp)
	utils.Check(err)
	strs := utils.StrArrayFromBytes(dat, "\n\n")
	ps := strArrayToPassports(strs)
	c1, rfps := countPassportsWithRequiredFields(ps)
	fmt.Println(c1)
	c2 := countValidPassports(rfps)
	fmt.Println(c2)
}
