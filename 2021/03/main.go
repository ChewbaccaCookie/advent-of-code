package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := shared.ReadFileRows("input.txt")

	inputColumn := [][]int{}
	for i := 0; i < len(input[0]); i++ {
		inputColumn = append(inputColumn, []int{})
	}

	inputRow := [][]int{}
	for i, row := range input {
		rowItems := strings.Split(row, "")
		inputRow = append(inputRow, []int{})

		for ii, item := range rowItems {
			inputColumn[ii] = append(inputColumn[ii], shared.Str2Int(item))
			inputRow[i] = append(inputRow[i], shared.Str2Int(item))
		}
	}
	gamma := ""
	epsilon := ""
	for _, column := range inputColumn {
		numNull := 0
		numOne := 0
		for _, item := range column {
			if item == 0 {
				numNull++
			} else {
				numOne++
			}
		}
		if numOne > numNull {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println(gammaDec * epsilonDec)

	// Part 2
	o2 := filterRows(0, len(input[0]), inputRow, true)
	co2 := filterRows(0, len(input[0]), inputRow, false)

	o2Dec, _ := strconv.ParseInt(shared.IntArr2String(o2[0]), 2, 64)
	co2Dec, _ := strconv.ParseInt(shared.IntArr2String(co2[0]), 2, 64)

	fmt.Println(o2Dec * co2Dec)

}

func filterRows(index, maxLen int, inputRows [][]int, keepMost bool) [][]int {
	if index == maxLen || len(inputRows) == 1 {
		return inputRows
	}

	numNull := 0
	numOne := 0
	for _, row := range inputRows {
		if row[index] == 1 {
			numOne++
		} else {
			numNull++
		}
	}
	if numNull == numOne {
		numOne++
	}

	newArr := [][]int{}
	for _, row := range inputRows {
		if (keepMost && (numNull > numOne && row[index] == 0 || numOne > numNull && row[index] == 1)) || (!keepMost && (numNull > numOne && row[index] == 1 || numOne > numNull && row[index] == 0)) {
			newArr = append(newArr, row)
		}
	}
	return filterRows(index+1, maxLen, newArr, keepMost)

}
