package mechs

import (
	"terminal-tetris2/shape"
	"terminal-tetris2/utils"
)

// --------------------------------------------------------------------------------------|

// ApplyAction получает действие от пользователя, если оно есть
func ApplyAction(control *utils.Controls, gameVar *utils.GameVar) {
	control.MutexEvent.Lock()

	if control.NewEvent {
		processEvent(control, gameVar)
		control.NewEvent = false
	}
	control.MutexEvent.Unlock()
}

// --------------------------------------------------------------------------------------|

// processEvent обрабатывает нажатую клавишу
func processEvent(control *utils.Controls, gameVar *utils.GameVar) {
	if control.Ev.Key != 0 {
		switch control.Ev.Key {
		case utils.KEY_LEFT:
			gameVar.Object.Shape.Move(-1, 0)
			if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
				gameVar.Object.Shape.Move(1, 0)
			}
		case utils.KEY_RIGHT:
			gameVar.Object.Shape.Move(1, 0)
			if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
				gameVar.Object.Shape.Move(-1, 0)
			}
		case utils.KEY_ROTATE:
			gameVar.Object.Shape.Rotate(shape.ROTATE_FORWARD)
			if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
				gameVar.Object.Shape.Rotate(shape.ROTATE_BACKWARDS)
			}
		case utils.KEY_DROP:
			for !utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
				gameVar.Object.Shape.Move(0, 1)
			}
			gameVar.Object.Shape.Move(0, -1)
		case utils.KEY_SPEEDUP:
			// компенсируем LastLine
			gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, 1).Y
			if utils.CheckCollision(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks()) {
				// компенсируем LastLine
				gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, -1).Y
			}
		case utils.KEY_QUIT:
			gameVar.Quit = true
		}

	} else if control.Ev.Ch != 0 {
		switch control.Ev.Ch {
		case utils.KEY_SHOW_NEXT:
			gameVar.ShowNS = !gameVar.ShowNS
		case utils.KEY_HIDE_HINT:
			gameVar.HideHint = !gameVar.HideHint
		case utils.KEY_LEVEL_UP:
			gameVar.Stat.Level = min(gameVar.Stat.Level+1, utils.MAX_LEVEL)
		}
	}
}
