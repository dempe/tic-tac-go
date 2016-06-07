package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const gridSize int = 3

type Board struct {
	tiles [gridSize][gridSize]int
}

func main() {
	b := constructInitialBoard()
	playing := true

	for playing {
		printBoard(b)
		position, mark := readInput(b)

		err := placeMove(&b, position, mark)

		if err != nil {
			fmt.Println(err)
			break
		}

		playing = !isGameOver(b)

		fmt.Println()
	}
}

func isGameOver(b Board) bool {
	return checkRows(b) || checkDiagonals(b) || checkCols(b)
}

func checkCols(b Board) bool {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[0][i]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[1][i] && b.tiles[1][i] == b.tiles[2][i] {
			return true
		}
	}

	return false
}

func checkDiagonals(b Board) bool {
	topLeft := b.tiles[0][0]
	topRight := b.tiles[0][2]
	middle := b.tiles[1][1]
	bottomLeft := b.tiles[2][0]
	bottomRight := b.tiles[2][2]

	if middle == 0 {
		return false
	}

	return (topLeft == middle && middle == bottomRight) || (topRight == middle && middle == bottomLeft)
}

func checkRows(b Board) bool {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[i][0]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] {
			return true
		}
	}

	return false
}

func placeMove(b *Board, position []int, mark string) error {
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

func readInput(b Board) ([]int, string) {
	fmt.Println("Please input your move in the form:  row,col,type.  Example:  0,2,X")
	fmt.Println("Rows and columns go from 0 - 2, while types can be one of 'X' or 'O'")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	position, mark, err := parseInput(input)

	if err != nil {
		fmt.Println(err)
	}

	return position, mark
}

func parseInput(input string) ([]int, string, error) {
	arr := strings.Split(input, ",")

	if len(arr) < gridSize {
		return nil, "", errors.New("Unrecognized input")
	}

	row, rowErr := strconv.Atoi(arr[0])
	col, colErr := strconv.Atoi(arr[1])
	mark := strings.Replace(arr[2], "\n", "", 1)

	if rowErr != nil || colErr != nil {
		return nil, "", errors.New("Unrecognized input")
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, "", errors.New("Unrecognized input")
	}

	if mark != "X" && mark != "O" {
		return nil, "", errors.New("Unrecognized input" + mark)
	}

	return []int{row, col}, mark, nil
}

func constructInitialBoard() Board {
	return Board{[gridSize][gridSize]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}}
}

func printBoard(b Board) {
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
			fmt.Print(decodeValue(row[0]))
		case 4:
			fmt.Print(decodeValue(row[1]))
		case 7:
			fmt.Print(decodeValue(row[2]))
		default:
			fmt.Print(" _ ")
		}
	}

	fmt.Println("|")
}

func decodeValue(value int) string {
	switch value {
	case 0:
		return ""
	case 1:
		return "O"
	case 2:
		return "X"
	}

	return ""
}
