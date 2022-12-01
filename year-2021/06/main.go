package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"strings"
)

func main() {
	input := shared.ReadFile("input.txt")
	laternFishTimers := strings.Split(input, ",")

	laternFishes := [10]int64{}

	for _, timer := range laternFishTimers {
		timerInt := shared.Str2Int(timer)
		laternFishes[timerInt]++
	}

	fmt.Println("A", calcLaternFishes(80, laternFishes))
	fmt.Println("B", calcLaternFishes(256, laternFishes))

}

func calcLaternFishes(days int, laternFishes [10]int64) int64 {
	for day := 0; day < days; day++ {
		laternFishes[7] += laternFishes[0]
		laternFishes[9] += laternFishes[0]
		laternFishes[0] = 0
		for i := 0; i < len(laternFishes)-1; i++ {
			laternFishes[i] = laternFishes[i+1]
		}
		laternFishes[len(laternFishes)-1] = 0
	}

	count := int64(0)
	for _, v := range laternFishes {
		count += v
	}
	return count
}
