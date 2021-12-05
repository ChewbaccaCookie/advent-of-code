package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"math"
)

type Rule struct {
	x1       int
	y1       int
	x2       int
	y2       int
	dX       int
	dY       int
	diagonal bool
}

func main() {
	ruleStrings := shared.ReadFileRows("input.txt")

	var maxX int
	var maxY int
	rules := []Rule{}

	// Step 1: Predefine map size
	for _, rule := range ruleStrings {
		match := shared.RegexMatch(`^(\d*),(\d*)\s->\s(\d*),(\d*)$`, rule)
		new := Rule{
			x1: shared.Str2Int(match[1]),
			y1: shared.Str2Int(match[2]),
			x2: shared.Str2Int(match[3]),
			y2: shared.Str2Int(match[4]),
		}

		new.dX = int(math.Abs(float64(new.x1 - new.x2)))
		new.dY = int(math.Abs(float64(new.y1 - new.y2)))

		if new.x1 > maxX {
			maxX = new.x1
		}
		if new.x2 > maxX {
			maxX = new.x2
		}
		if new.y1 > maxY {
			maxY = new.y1
		}
		if new.y2 > maxX {
			maxY = new.y2
		}

		// Check rules
		if !(new.x1 == new.x2 || new.y1 == new.y2) {
			new.diagonal = true
		}

		rules = append(rules, new)

	}

	// Step 2: Generate Empty map
	ventMapA := [][]int{}
	ventMapB := [][]int{}
	for i := 0; i <= maxX; i++ {
		ventMapA = append(ventMapA, []int{})
		ventMapB = append(ventMapB, []int{})
		for ii := 0; ii <= maxY; ii++ {
			ventMapA[i] = append(ventMapA[i], 0)
			ventMapB[i] = append(ventMapB[i], 0)
		}
	}

	// Step 3: Generate Vector in vent map
	for _, rule := range rules {
		if rule.diagonal {
			y := rule.y1
			x := rule.x1
			for i := 0; i <= rule.dX; i++ {
				ventMapB[x][y]++
				if rule.y1 < rule.y2 {
					y++
				} else {
					y--
				}
				if rule.x1 < rule.x2 {
					x++
				} else {
					x--
				}
			}
		} else {
			x := rule.x1
			for i := 0; i <= rule.dX; i++ {
				y := rule.y1
				for ii := 0; ii <= rule.dY; ii++ {
					ventMapA[x][y]++
					ventMapB[x][y]++

					if rule.y1 < rule.y2 {
						y++
					} else {
						y--
					}
				}

				if rule.x1 < rule.x2 {
					x++
				} else {
					x--
				}
			}
		}
	}

	// Step 4: Count Dangerous tiles
	dangerousA := 0
	for _, row := range ventMapA {
		for _, cell := range row {
			if cell >= 2 {
				dangerousA++
			}
		}
	}

	dangerousB := 0
	for _, row := range ventMapB {
		for _, cell := range row {
			if cell >= 2 {
				dangerousB++
			}
		}
	}
	fmt.Println(dangerousA)
	fmt.Println(dangerousB)
}
