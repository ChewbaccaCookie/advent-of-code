package shared

import "strconv"

func Str2Int(strVar string) int {
	i, err := strconv.Atoi(strVar)
	if err != nil {
		panic(err)
	}
	return i
}

func Int2Str(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func IntArr2String(intArr []int) (out string) {
	for _, item := range intArr {
		out += Int2Str(item)
	}
	return out
}

func StringArr2IntStr(strArr []string) (out []int) {
	for _, item := range strArr {
		out = append(out, Str2Int(item))
	}
	return out
}
