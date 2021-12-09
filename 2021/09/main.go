package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"strings"
)

func main() {
	input := shared.ReadFileRows("input.txt")
	locationMap := [][]int{}
	// Prefill map
	for i, row := range input {
		locationMap = append(locationMap, []int{})
		for _, cell := range strings.Split(row, "") {
			locationMap[i] = append(locationMap[i], shared.Str2Int(cell))
		}
	}

	// Key for this map is row-column
	// Int is the height of the tile
	dangerZones := map[string]int{}

	// Do for each item
	for x, row := range locationMap {
		for y := range row {
			searchLowest(x, y, 9, &locationMap, &dangerZones)
		}
	}

	// Calc risk
	risk := 0
	for _, v := range dangerZones {
		risk += v + 1
	}

	fmt.Println(risk)

}

func searchLowest(row, cell, lastVal int, locationMap *[][]int, dangerZones *map[string]int) (value int, foundAlready bool) {
	rowSize := len(*locationMap)
	columnSize := len((*locationMap)[0])

	// Abbruchbedingung
	if row >= rowSize || row < 0 || cell < 0 || cell >= columnSize {
		return -1, false
	}

	value = (*locationMap)[row][cell]
	if value >= lastVal {
		return value, false
	}

	lowestVal := value
	lowestPos := ""
	for i := 0; i < 4; i++ {
		y, x := getNextCoorinates(i)
		val, found := searchLowest(row+y, cell+x, value, locationMap, dangerZones)
		if found {
			return val, true
		}
		if val != -1 && val < lowestVal {
			lowestVal = val
			lowestPos = fmt.Sprintf("%d-%d", row+y, cell+x)
		}
	}
	if lowestPos != "" {
		(*dangerZones)[lowestPos] = lowestVal
		return lowestVal, true
	}
	return lowestVal, false
}

func getNextCoorinates(iteration int) (row, cell int) {
	switch iteration {
	case 0:
		return 0, 1
	case 1:
		return -1, 0
	case 2:
		return 0, -1
	case 3:
		return 1, 0
	}
	return 0, 0
}
