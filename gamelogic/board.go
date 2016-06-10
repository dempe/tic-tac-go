package gamelogic

import (
	"errors"
	"fmt"
)

const gridSize int = 3

type Board struct {
	tiles [gridSize][gridSize]int
}

type Winner struct {
	Mark         string
	Undetermined bool
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

func (b *Board) Copy() *Board {
	newBoard := NewBoard()

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			newBoard.tiles[i][j] = b.tiles[i][j]
		}
	}

	return &newBoard
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

func (b *Board) GetWinningPlayer() *Winner {
	rowVictory := b.getRowVictory()
	diagonalVictory := b.getDiagonalVictory()
	columnVictory := b.getColumnVictory()

	if rowVictory != nil {
		return rowVictory
	} else if columnVictory != nil {
		return columnVictory
	} else if diagonalVictory != nil {
		return diagonalVictory
	} else if b.isFilled() {
		return &Winner{"", false}
	}

	return &Winner{"", true}
}

func (b *Board) getColumnVictory() *Winner {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[0][i]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[1][i] && b.tiles[1][i] == b.tiles[2][i] {
			return &Winner{DecodeValue(firstCell), false}
		}
	}

	return &Winner{"", true}
}

func (b *Board) getDiagonalVictory() *Winner {
	topLeft := b.tiles[0][0]
	topRight := b.tiles[0][2]
	middle := b.tiles[1][1]
	bottomLeft := b.tiles[2][0]
	bottomRight := b.tiles[2][2]

	if middle == 0 {
		return &Winner{"", true}
	}

	if topLeft == middle && middle == bottomRight {
		return &Winner{DecodeValue(topLeft), false}
	} else if topRight == middle && middle == bottomLeft {
		return &Winner{DecodeValue(topRight), false}
	}

	return &Winner{"", true}
}

func (b *Board) getRowVictory() *Winner {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[i][0]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] {
			return &Winner{DecodeValue(firstCell), false}
		}
	}

	return nil
}

func (b *Board) isFilled() bool {
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if b.tiles[i][j] == 0 {
				return false
			}
		}
	}

	return true
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
