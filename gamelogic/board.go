package gamelogic

import (
	"errors"
	"fmt"
)

const gridSize int = 3

type Board struct {
	tiles [gridSize][gridSize]int
}

func NewBoard() Board {
	return Board{[gridSize][gridSize]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}}
}

func (b *Board) GetTiles() [3][3]int {
	return b.tiles
}

func (b *Board) PrintBoard() {
	fmt.Println("Here is the current board:")
	for i := 0; i < gridSize; i++ {
		fmt.Println()
		printRow(b.tiles[i])
	}
}

func printRow(row [gridSize]int) {
	for i := 0; i < 9; i++ {
		if i%gridSize == 0 {
			fmt.Print("|")
		}

		switch i {
		case 1:
			fmt.Print(DecodeValue(row[0]))
		case 4:
			fmt.Print(DecodeValue(row[1]))
		case 7:
			fmt.Print(DecodeValue(row[2]))
		default:
			fmt.Print(" _ ")
		}
	}

	fmt.Println("|")
}

func (b *Board) PlaceMove(position []int, mark string) error {
	if b.tiles[position[0]][position[1]] != 0 {
		return errors.New("Position already occupided")
	}

	switch mark {
	case "O":
		b.tiles[position[0]][position[1]] = 1
	case "X":
		b.tiles[position[0]][position[1]] = 2
	}

	return nil
}

func (b *Board) GetWinningPlayer() string {
	rowVictory := b.getRowVictory()
	diagonalVictory := b.getDiagonalVictory()
	columnVictory := b.getColumnVictory()

	if rowVictory != "" {
		fmt.Println(rowVictory + " wins!")
		return rowVictory
	} else if columnVictory != "" {
		fmt.Println(columnVictory + " wins!")
		return columnVictory
	} else if diagonalVictory != "" {
		fmt.Println(diagonalVictory + " wins!")
		return diagonalVictory
	}

	return ""
}

func (b *Board) getColumnVictory() string {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[0][i]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[1][i] && b.tiles[1][i] == b.tiles[2][i] {
			return DecodeValue(firstCell)
		}
	}

	return ""
}

func (b *Board) getDiagonalVictory() string {
	topLeft := b.tiles[0][0]
	topRight := b.tiles[0][2]
	middle := b.tiles[1][1]
	bottomLeft := b.tiles[2][0]
	bottomRight := b.tiles[2][2]

	if middle == 0 {
		return ""
	}

	if topLeft == middle && middle == bottomRight {
		return DecodeValue(topLeft)
	} else if topRight == middle && middle == bottomLeft {
		return DecodeValue(topRight)
	}

	return ""
}

func (b *Board) getRowVictory() string {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[i][0]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] {
			return DecodeValue(firstCell)
		}
	}

	return ""
}

func DecodeValue(value int) string {
	switch value {
	case 0:
		return " "
	case 1:
		return "O"
	case 2:
		return "X"
	}

	return ""
}
