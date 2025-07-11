package utils

import (
	"terminal-tetris2/shape"
)

// ------------------------------------------------------------------------------------|

// инициализация игрового поля, фигуры и старых кирпичиков
func Initialize(level int) GameVar {
	var gameVar GameVar

	gameVar.Object.OldBricks = make([]shape.Brick, 0)
	gameVar.Object.Shape = shape.GetNewShape()
	gameVar.Object.NextShape = shape.GetNewShape()

	gameVar.Object.Shape.Move(4, -1) // перемещаем фигуру в начальное положение

	gameVar.Stat.LastLine = BASE_SCORE
	gameVar.Stat.Score = 0
	gameVar.Stat.Level = level

	gameVar.Timer.Fell = TIMER_OFF
	gameVar.Timer.Clear = TIMER_OFF

	return gameVar
}
