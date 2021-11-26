package shared

import "strconv"

func Str2Int(strVar string) int {
	i, _ := strconv.Atoi(strVar)
	return i
}
