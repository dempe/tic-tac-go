package ai

import (
	"container/list"
	"errors"
	"fmt"

	"github.com/dempe/tictacgo/gamelogic"
)

type Score struct {
	Score         int
	NotDetermined bool
}

type GameState struct {
	board     gamelogic.Board
	turn      int
	subStates *list.List
}

func NewGameState(b gamelogic.Board) (*GameState, error) {
	turn := whoseTurn(b)
	moves := CalculatePossibleMoves(b)
	subStates := list.New()

	for e := moves.Front(); e != nil; e = e.Next() {
		newBoard := b.Copy()

		value, ok := e.Value.([]int)

		if !ok {
			return nil, errors.New("Expected type []int")
		}

		b.PlaceMove(value, gamelogic.DecodeValue(turn))
		subState, _ := NewGameState(*newBoard)
		subStates.PushBack(subState)
	}

	return &GameState{b, turn, subStates}, nil
}

func whoseTurn(b gamelogic.Board) int {
	var xcount, ycount int
	tiles := b.GetTiles()

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 1 {
				ycount++
			} else if tiles[i][j] == 2 {
				xcount++
			}
		}
	}

	// X always goes first
	if xcount == ycount {
		return 2
	}

	return 1
}

func CalculateScore(b gamelogic.Board, mark string) (Score, error) {
	if mark != "X" && mark != "O" {
		return Score{0, true}, fmt.Errorf("Unrecognized mark, %s.  Must be X or O", mark)
	}

	winner := b.GetWinningPlayer()

	if winner.Undetermined {
		return Score{0, false}, nil
	} else if mark == winner.Mark {
		return Score{1, false}, nil
	}

	return Score{-1, false}, nil
}

func CalculatePossibleMoves(b gamelogic.Board) *list.List {
	positions := list.New()
	tiles := b.GetTiles()

	for i := 0; i < 3; i++ {
		for j := 0; j < len(tiles[i]); j++ {
			if tiles[i][j] == 0 {
				positions.PushBack([2]int{i, j})
			}
		}
	}

	return positions
}
