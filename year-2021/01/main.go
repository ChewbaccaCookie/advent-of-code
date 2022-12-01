package main

import (
	"advent-of-code-2020/shared"
	"fmt"
)

func main() {
	lines := shared.ReadFileRows("input.txt")
	increasedA := 1
	increasedB := 0
	old := -1
	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1] > lines[i] {
			increasedA++
		}
		if i == len(lines)-2 {
			break
		}

		new := (shared.Str2Int(lines[i]) + shared.Str2Int(lines[i+1]) + shared.Str2Int(lines[i+2]))
		if old != -1 && new > old {
			increasedB++
		}
		old = new
	}
	fmt.Println(increasedA)
	fmt.Println(increasedB)
}
