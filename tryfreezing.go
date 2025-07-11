package mechs

import (
	"terminal-tetris2/shape"
	"terminal-tetris2/utils"
	"time"
)

// ---------------------------------------------------------------------------------------------------------|

// TryFreezing проверяет упал ли фигура
func TryFreezing(gameVar *utils.GameVar) {
	freez := false

	gameVar.Object.Shape.Move(0, 1)
	if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
		freez = true // нужно заморозить
	}
	gameVar.Object.Shape.Move(0, -1)

	if freez {
		freezing(gameVar)
	}
}

// ---------------------------------------------------------------------------------------------------------|

func freezing(gameVar *utils.GameVar) {
	generateNewShape(gameVar)
	gameVar.Stat.FullLines = FindFullLines(gameVar.Object.OldBricks) // ищем полные линии в игровом поле
	countStats(gameVar)

	if len(gameVar.Stat.FullLines) > 0 {
		// компенсируем Lastline
		gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(4, -2).Y // фигура выходит из центра и прячется вне поля
		gameVar.CurrentTime = time.Now().UnixMilli()
		gameVar.Timer.Clear = gameVar.CurrentTime + utils.CLEAR_TIME*utils.T_QUANT
	} else {
		// компенсируем Lastline
		gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(4, 1).Y // фигура выходит из центра
		if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
			gameVar.Quit = true // если новая фигура не может выйти то выход из игры
		}
		gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, -1).Y
	}
}

// ---------------------------------------------------------------------------------------------------------|

func generateNewShape(gameVar *utils.GameVar) {
	gameVar.Object.OldBricks = append(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()...)
	shape.SortBricksByY(gameVar.Object.OldBricks)
	gameVar.Object.Shape = gameVar.Object.NextShape
	gameVar.Object.NextShape = shape.GetNewShape()
}

// ---------------------------------------------------------------------------------------------------------|

func countStats(gameVar *utils.GameVar) {
	// обновляем количество удаленных линий, делим на 2, т.к. каждая полная линия состоит из 2-х кирпичиков
	gameVar.Stat.DelLines += len(gameVar.Stat.FullLines) / 2

	// повышаем уровень, если удалено 4 строки одновременно
	if len(gameVar.Stat.FullLines)/2 == utils.DEL_LINES_LEVEL_UP {
		gameVar.Stat.Level++
	}
	// обновляем количество очков за последнюю линию, но не меньше 0
	gameVar.Stat.LastLine = max(gameVar.Stat.LastLine, 0)
	gameVar.Stat.Score += gameVar.Stat.LastLine + gameVar.Stat.Level*utils.LEVEL_BONUS
	if !gameVar.ShowNS {
		gameVar.Stat.Score += utils.FORESIGHT_BONUS
	}
	gameVar.Stat.LastLine = utils.BASE_SCORE
}
