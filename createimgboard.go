package rendering

import (
	"fmt"
	"terminal-tetris2/utils"
)

// --------------------------------------------------------------------------------------|

func CreateImgBoard(topScores []utils.ScoreEntry, markedIndex int) Img {
	board := make(Img, 0, 11)

	board = append(board, "ИМЯ        УРОВЕНЬ     СЧЕТ")

	for i, entry := range topScores {
		nameStr, ok1 := entry.Name.Value.(string)
		levelInt, ok2 := entry.Level.Value.(int)
		scoreInt, ok3 := entry.Score.Value.(int)

		if !ok1 || !ok2 || !ok3 {
			continue
		}

		line := fmt.Sprintf("%-10s  %-9d  %-2d", nameStr, levelInt, scoreInt)

		if i == markedIndex {
			line += " **"
		}

		board = append(board, line)
	}

	return board
}

// --------------------------------------------------------------------------------------|

func CreateImgNextGame() Img {
	nextGame := make(Img, 0, 1)

	nextGame = append(nextGame, "ЕЩЕ ПАРТИЮ? (ДА/НЕТ) -")

	return nextGame
}
