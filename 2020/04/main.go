package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type PasswortValidation struct {
	key        string
	validation func(val string) bool
}

func main() {
	start := time.Now()
	input := shared.ReadFile("input.txt")
	passports := strings.Split(input, "\n\n")
	properties := []PasswortValidation{
		{
			key: "byr:",
			validation: func(val string) bool {
				return shared.Str2Int(val) >= 1920 && shared.Str2Int(val) <= 2002
			},
		}, {
			key: "iyr:",
			validation: func(val string) bool {
				return shared.Str2Int(val) >= 2010 && shared.Str2Int(val) <= 2020
			},
		},
		{
			key: "eyr:",
			validation: func(val string) bool {
				return shared.Str2Int(val) >= 2020 && shared.Str2Int(val) <= 2030
			},
		},
		{
			key: "hgt:",
			validation: func(val string) bool {
				re := regexp.MustCompile(`(\d*)(cm|in)`)
				str := re.FindStringSubmatch(val)
				if len(str) < 3 {
					return false
				}
				if str[2] == "cm" {
					return shared.Str2Int(str[1]) >= 150 && shared.Str2Int(str[1]) <= 193
				} else {
					return shared.Str2Int(str[1]) >= 59 && shared.Str2Int(str[1]) <= 76
				}
			},
		},
		{
			key: "hcl:",
			validation: func(val string) bool {
				v, _ := regexp.MatchString("#([a-f]|[0-9]){6}$", val)
				return v
			},
		},
		{
			key: "ecl:",
			validation: func(val string) bool {
				allowed := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
				return allowed[val]
			},
		},
		{
			key: "pid:",
			validation: func(val string) bool {
				v, _ := regexp.MatchString("^[0-9]{9}$", val)
				return v
			},
		},
	}

	validA := 0
	validB := 0
	for _, passport := range passports {
		p := strings.ReplaceAll(passport, "\n", " ")
		v1 := true
		v2 := true
		for _, property := range properties {
			if strings.Contains(p, property.key) {
				re := regexp.MustCompile(fmt.Sprintf(`%s(\S*)`, property.key))
				str := re.FindStringSubmatch(p)
				if !property.validation(str[1]) {
					v2 = false
				}
			} else {
				v1 = false
				v2 = false
				break
			}
		}
		if v1 {
			validA++
		}
		if v2 {
			validB++
		}
	}

	fmt.Println(validA)
	fmt.Println(validB)
	fmt.Println(time.Since(start))
}
