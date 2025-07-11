package mechs

import (
	"terminal-tetris2/utils"
	"time"
)

// ----------------------------------------------------------------------------------------------------------------------------|

// TryMoveDown применяет действие к фигуре, двигая ее вниз по игровому полю
func TryMoveDown(gameVar *utils.GameVar) {
	// компенсируем LastLine
	gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, 1).Y

	LevelFellSpeed := utils.FELL_TIME - int64(gameVar.Stat.Level+1)*utils.FELL_TIME_INCREASER

	if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
		// компенсируем LastLine
		gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, -1).Y

		if gameVar.Timer.Fell == utils.TIMER_OFF {
			gameVar.CurrentTime = time.Now().UnixMilli()
			gameVar.Timer.Fell = gameVar.CurrentTime + LevelFellSpeed*utils.T_QUANT
		}
	}
}
