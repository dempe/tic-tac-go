package ai

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dempe/tictacgo/gamelogic"
	"github.com/dempe/tictacgo/minimax"
)

func ComputerMove(b gamelogic.Board, aiType, computerMark string) []int {
	fmt.Println("Computer's turn!")

	switch aiType {
	case "random":
		return moveRandomly(b)
	case "minimax":
		pos, err := moveUsingMiniMax(b, computerMark)
		if err != nil {
			fmt.Println(err)
		}
		return pos
	default:
		return moveRandomly(b)
	}
}

func findBoardDifference(original, modified gamelogic.Board, computerMark int) ([]int, error) {
	winner := modified.GetWinningPlayer()
	oppVictory := !winner.Undetermined && gamelogic.EncodeValue(winner.Mark) != computerMark

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if original.GetTiles()[i][j] != modified.GetTiles()[i][j] {
				if oppVictory && modified.GetTiles()[i][j] != computerMark {
					return []int{i, j}, nil
				} else if modified.GetTiles()[i][j] == computerMark {
					return []int{i, j}, nil
				}
			}
		}
	}

	return []int{}, errors.New("Found no difference")
}

func moveUsingMiniMax(b gamelogic.Board, computerMark string) ([]int, error) {
	state, err := minimax.NewGameState(b)

	if err != nil {
		fmt.Println(err)
	}

	gameStateScore, err := state.MiniMax(computerMark)

	if err != nil {
		fmt.Println(err)
	}

	pos, err := findBoardDifference(b, gameStateScore.Board, gamelogic.EncodeValue(computerMark))

	if err != nil {
		fmt.Println(err)
	}

	return pos, nil
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
