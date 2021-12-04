package main

import (
	"advent-of-code-2020/shared"
	"fmt"
	"strings"
)

type Board struct {
	boardNum    int
	drawnNumber int
	sum         int
}

func main() {
	rows := shared.ReadFileRows("input.txt")
	bingoNums := shared.StringArr2IntStr(strings.Split(rows[0], ","))

	bingoPieces := [][]int{{}}
	drawnFromPiece := map[int]int{}
	winningColumns := map[int]map[int]int{}
	winningRows := map[int]map[int]int{}
	boardsWon := []Board{}

	piece := 0
	for _, row := range rows[2:] {
		if row == "" {
			bingoPieces = append(bingoPieces, []int{})
			piece++
			continue
		}
		for _, v := range strings.Split(row, " ") {
			if v == "" {
				continue
			}
			bingoPieces[piece] = append(bingoPieces[piece], shared.Str2Int(v))
		}
	}

	for _, num := range bingoNums {
		for i, piece := range bingoPieces {
			for ii, pieceNum := range piece {
				if num == pieceNum {
					drawnFromPiece[i] += num
					if _, ok := winningRows[i]; !ok {
						winningRows[i] = map[int]int{}
					}
					if _, ok := winningColumns[i]; !ok {
						winningColumns[i] = map[int]int{}
					}

					winningRows[i][ii/5]++
					winningColumns[i][ii%5]++

					break
				}

			}
		}
		// Check if someone has won by column
		for pieceNum, columns := range winningColumns {
			for _, v := range columns {
				if v == 5 {
					sum := 0
					for _, piece := range bingoPieces[pieceNum] {
						sum += piece
					}
					boardsWon = append(boardsWon, Board{
						boardNum:    pieceNum,
						drawnNumber: num,
						sum:         sum - drawnFromPiece[pieceNum],
					})
					bingoPieces[pieceNum] = []int{}
					winningRows[pieceNum] = map[int]int{}
					winningColumns[pieceNum] = map[int]int{}
					break
				}
			}

		}

		if len(boardsWon) == len(bingoPieces) {
			break
		}

		for pieceNum, rows := range winningRows {
			for _, v := range rows {
				if v == 5 {
					sum := 0
					for _, piece := range bingoPieces[pieceNum] {
						sum += piece
					}
					boardsWon = append(boardsWon, Board{
						boardNum:    pieceNum,
						drawnNumber: num,
						sum:         sum - drawnFromPiece[pieceNum],
					})
					bingoPieces[pieceNum] = []int{}
					winningRows[pieceNum] = map[int]int{}
					winningColumns[pieceNum] = map[int]int{}
					break
				}

			}

		}
		if len(boardsWon) == len(bingoPieces) {
			break
		}

	}
	best := boardsWon[0]
	worst := boardsWon[len(boardsWon)-1]

	fmt.Println(best.sum * best.drawnNumber)
	fmt.Println(worst.sum * worst.drawnNumber)
}
