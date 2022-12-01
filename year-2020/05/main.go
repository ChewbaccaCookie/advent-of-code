package main

import (
	"advent-of-code-2020/shared"
	"fmt"
)

type DetailedSeat struct {
	raw    string
	row    int
	column int
	seatID int
}

func main() {
	seats := shared.ReadFileRows("input.txt")
	//seats := []string{"FBFBBFFRLR"}
	var detailedSeats []DetailedSeat
	var highestSeat DetailedSeat
	var seatMap [128][8]bool
	for _, seat := range seats {
		row := 0
		column := 0
		step := 64
		for i := 0; i < 7; i++ {
			if seat[i] == 'B' {
				row += step
			}
			step /= 2
		}
		step = 4
		for i := 7; i < 10; i++ {
			if seat[i] == 'R' {
				column += step
			}
			step /= 2
		}

		s := DetailedSeat{
			raw:    seat,
			row:    row,
			column: column,
			seatID: (row * 8) + column,
		}
		detailedSeats = append(detailedSeats, s)
		if s.seatID > highestSeat.seatID {
			highestSeat = s
		}

		seatMap[row][column] = true
	}
	emptyRow := -1
	emptyColumn := -1
	for i, row := range seatMap {
		//remove empty
		empty := 0
		for _, column := range row {
			if !column {
				empty++
			}
		}

		if empty > 2 {
			continue
		}
		for ii, column := range row {
			if !column {
				emptyRow = i
				emptyColumn = ii
				break
			}
		}
		if emptyRow != -1 {
			break
		}
	}

	fmt.Println(highestSeat)
	fmt.Println(DetailedSeat{
		row:    emptyRow,
		column: emptyColumn,
		seatID: (emptyRow * 8) + emptyColumn,
	})
}
