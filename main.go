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

	"github.com/dempe/tictacgo/gamelogic"
)

func main() {
	b := gamelogic.NewBoard()
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
			b.PrintBoard()
			err := b.PlaceMove(readInput(), playerMark)

			if err != nil {
				fmt.Println(err)
				break
			}

			playerTurn = false
		} else {
			b.PlaceMove(computerMove(b, "random"), computerMark)
			playerTurn = true
		}

		winningPlayer := b.GetWinningPlayer()
		playing = winningPlayer == ""

		fmt.Println()
	}
}

func computerMove(b gamelogic.Board, aiType string) []int {
	fmt.Println("Computer's turn!")

	switch aiType {
	case "random":
		return moveRandomly(b)
	default:
		return moveRandomly(b)
	}
}

func moveRandomly(b gamelogic.Board) []int {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	row := rand.Intn(3)
	col := rand.Intn(3)

	tiles := b.GetTiles()

	for tiles[row][col] != 0 {
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
