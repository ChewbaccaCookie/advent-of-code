package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	input := shared.ReadFile("input.txt")
	positions := shared.StringArr2IntStr(strings.Split(input, ","))
	sort.Ints(positions)

	min := positions[0]
	max := positions[len(positions)-1]
	bestCountA := 0
	bestCountDiffA := 0
	bestCountB := 0
	bestCountDiffB := 0
	for i := min; i < max; i++ {
		countA := 0
		countB := 0
		for _, v := range positions {
			diff := int(math.Abs(float64(v - i)))
			countA += diff
			for ii := 1; ii <= diff; ii++ {
				countB += ii
			}
		}

		if bestCountDiffA == 0 || countA < bestCountDiffA {
			bestCountDiffA = countA
			bestCountA = i
		}
		if bestCountDiffB == 0 || countB < bestCountDiffB {
			bestCountDiffB = countB
			bestCountB = i
		}
	}

	fmt.Println("A", bestCountA, bestCountDiffA)
	fmt.Println("B", bestCountB, bestCountDiffB)
}
