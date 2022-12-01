package main

import (
	"advent-of-code-2020/shared"
	"fmt"
)

func main() {
	lines := shared.ReadFileRows("input.txt")
	depthA := 0
	depthB := 0
	horizontalA := 0
	horizontalB := 0
	aim := 0
	for _, line := range lines {
		str := shared.RegexMatch(`^([a-z]*)\s(\d)$`, line)
		command := str[1]
		distance := shared.Str2Int(str[2])

		switch command {
		case "forward":
			horizontalA += distance
			horizontalB += distance
			depthB += aim * distance
		case "up":
			depthA -= distance
			aim -= distance
		case "down":
			depthA += distance
			aim += distance
		}
	}

	fmt.Println(depthA * horizontalA)
	fmt.Println(depthB * horizontalB)
}
