package ai

import (
	"fmt"

	"github.com/dempe/tictacgo/gamelogic"
)

type Score struct {
	score         int
	notDetermined bool
}

func CalculateScore(b gamelogic.Board, mark string) (Score, error) {
	if mark != "X" && mark != "O" {
		return Score{0, true}, fmt.Errorf("Unrecognized mark, %s.  Must be X or O", mark)
	}

	winner := b.GetWinningPlayer()

	if winner == "" {
		return Score{0, false}, nil
	} else if mark == winner {
		return Score{1, false}, nil
	}

	return Score{-1, false}, nil
}

func (s *Score) GetScore() int {
	return s.score
}
