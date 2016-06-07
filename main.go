package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const gridSize int = 3

type Board struct {
	tiles [gridSize][gridSize]int
}

func main() {
	b := constructInitialBoard()
	playing := true
	playerTurn := playerGoesFirst()
	var computerMark, playerMark string

	if !playerTurn {
		computerMark = "X"
		playerMark = "O"
	} else {
		computerMark = "O"
		playerMark = "X"
	}

	for playing {
		if playerTurn {
			fmt.Println("Your turn!")
			printBoard(b)
			err := placeMove(&b, readInput(), playerMark)

			if err != nil {
				fmt.Println(err)
				break
			}

			playerTurn = false
		} else {
			placeMove(&b, computerMove(b, "random"), computerMark)
			playerTurn = true
		}

		winningPlayer := getWinningPlayer(b)
		playing = winningPlayer == ""

		fmt.Println()
	}
}

func computerMove(b Board, aiType string) []int {
	fmt.Println("Computer's turn!")

	switch aiType {
	case "random":
		return moveRandomly(b)
	default:
		return moveRandomly(b)
	}
}

func moveRandomly(b Board) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	row := rand.Intn(3)
	col := rand.Intn(3)

	for b.tiles[row][col] != 0 {
		row = rand.Intn(3)
		col = rand.Intn(3)
	}

	return []int{row, col}
}

func playerGoesFirst() bool {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	return rand.Intn(100)%2 == 0
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

func readInput() []int {
	fmt.Println("Please input your move in the form:  row,col.  Example:  0,2")

	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	position, err := parseInput(input)

	if err != nil {
		fmt.Println(err)
	}

	return position
}

func parseInput(input string) ([]int, error) {
	arr := strings.Split(input, ",")

	if len(arr) < 2 {
		return nil, errors.New("Unrecognized input - not enough values")
	}

	row, rowErr := strconv.Atoi(arr[0])
	col, colErr := strconv.Atoi(strings.Replace(arr[1], "\n", "", 1))

	if rowErr != nil {
		return nil, rowErr
	}

	if colErr != nil {
		return nil, colErr
	}

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return nil, errors.New("Unrecognized input - value must be in range [0, 2]")
	}

	return []int{row, col}, nil
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
		return " "
	case 1:
		return "O"
	case 2:
		return "X"
	}

	return ""
}

func getWinningPlayer(b Board) string {
	rowVictory := getRowVictory(b)
	diagonalVictory := getDiagonalVictory(b)
	columnVictory := getColumnVictory(b)

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

func getColumnVictory(b Board) string {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[0][i]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[1][i] && b.tiles[1][i] == b.tiles[2][i] {
			return decodeValue(firstCell)
		}
	}

	return ""
}

func getDiagonalVictory(b Board) string {
	topLeft := b.tiles[0][0]
	topRight := b.tiles[0][2]
	middle := b.tiles[1][1]
	bottomLeft := b.tiles[2][0]
	bottomRight := b.tiles[2][2]

	if middle == 0 {
		return ""
	}

	if topLeft == middle && middle == bottomRight {
		return decodeValue(topLeft)
	} else if topRight == middle && middle == bottomLeft {
		return decodeValue(topRight)
	}

	return ""
}

func getRowVictory(b Board) string {
	for i := 0; i < gridSize; i++ {
		firstCell := b.tiles[i][0]

		if firstCell == 0 {
			continue
		}

		if firstCell == b.tiles[i][1] && b.tiles[i][1] == b.tiles[i][2] {
			return decodeValue(firstCell)
		}
	}

	return ""
}
