package shared

import "strconv"

func Str2Int(strVar string) int {
	i, err := strconv.Atoi(strVar)
	if err != nil {
		panic(err)
	}
	return i
}
