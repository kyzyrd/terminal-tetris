package gameparts

import (
	"terminal-tetris2/mechs"
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"
	"time"

	"github.com/nsf/termbox-go"
)

// ----------------------------------------------------------------------------------------------------|

func Game(control *utils.Controls, canvas *rendering.Canvas, level int) *utils.ScoreEntry {
	gameVar := utils.Initialize(level)

	for !gameVar.Quit {
		gameVar.CurrentTime = time.Now().UnixMilli()

		fellTimerExpired := gameVar.CurrentTime > gameVar.Timer.Fell
		stepTimerExpired := gameVar.CurrentTime > gameVar.Timer.Step
		clearTimerExpired := gameVar.CurrentTime > gameVar.Timer.Clear
		clearTimerOFF := gameVar.Timer.Clear == utils.TIMER_OFF

		levelSpeed := utils.BASE_SPEED - (int64(gameVar.Stat.Level+1) * utils.SPEED_INCREASER)

		if clearTimerOFF {
			mechs.ApplyAction(control, &gameVar)

			if stepTimerExpired {
				mechs.TryMoveDown(&gameVar)
				gameVar.CurrentTime = time.Now().UnixMilli()
				gameVar.Timer.Step = gameVar.CurrentTime + levelSpeed*utils.T_QUANT
			}

			if fellTimerExpired {
				gameVar.Timer.Fell = utils.TIMER_OFF
				mechs.TryFreezing(&gameVar)
			}
		}

		if clearTimerExpired {
			gameVar.Timer.Clear = utils.TIMER_OFF
			mechs.ClearLines(&gameVar.Object.OldBricks, gameVar.Stat.FullLines)
			mechs.DropLines(gameVar.Object.OldBricks)

			// фигура падает вниз на 2 клетки, чтобы не зависала в воздухе
			gameVar.Stat.LastLine -= gameVar.Object.Shape.Move(0, 2).Y

			// обновляем текущее время
			gameVar.CurrentTime = time.Now().UnixMilli()
			gameVar.Timer.Step = gameVar.CurrentTime + levelSpeed*utils.T_QUANT
		}

		if gameVar.Timer.Clear != utils.TIMER_OFF {
			mechs.BlinkLines(gameVar.Object.OldBricks, gameVar.Stat.FullLines)
		}

		renderWindow(gameVar, canvas)

		time.Sleep(utils.TICK * utils.T_QUANT * time.Millisecond)
	}

	renderWindow(gameVar, canvas)

	return utils.CreateNewEntry(
		getUserName(control, canvas),
		gameVar.Stat.Level,
		gameVar.Stat.Score,
	)
}

// ----------------------------------------------------------------------------------------------------|

func renderWindow(gameVar utils.GameVar, canvas *rendering.Canvas) {
	canvas.Clear()

	newImg := rendering.CreateImgField(gameVar.Object.OldBricks, gameVar.Object.Shape.GetBricks())
	canvas.SetImage(newImg, 20, 0)
	newImg = rendering.CreateImgScore(gameVar.Stat.Score, gameVar.Stat.Level, gameVar.Stat.DelLines)
	canvas.SetImage(newImg, 0, 0)
	if gameVar.ShowNS {
		newImg = rendering.CreateImgNS(*gameVar.Object.NextShape)
		canvas.SetImage(newImg, 11, 10)
	}
	if !gameVar.HideHint {
		canvas.SetImage(rendering.CreateImgHint(), 46, 1)
	}

	if gameVar.Quit {
		canvas.SetImage(rendering.Img([]string{"ВАШЕ ИМЯ?"}), 0, 13)
	}

	canvas.Print()
}

// ----------------------------------------------------------------------------------------------------|

func getUserName(control *utils.Controls, canvas *rendering.Canvas) string {
	name := make([]rune, 0, 8)

	for {
		control.MutexEvent.Lock()
		if control.NewEvent {

			if control.Ev.Type == termbox.EventKey {
				switch {
				case control.Ev.Key == termbox.KeyEnter && len(name) > 0:
					control.MutexEvent.Unlock()
					return string(name)
				case control.Ev.Ch != 0 && len(name) < 8:
					name = append(name, control.Ev.Ch)
					renderNameInput(name, canvas)
				case (control.Ev.Key == termbox.KeyBackspace || control.Ev.Key == termbox.KeyBackspace2):
					if len(name) > 0 {
						name = name[:len(name)-1]
						renderNameInput(name, canvas)
					}
				}
			}

			control.NewEvent = false
		}
		control.MutexEvent.Unlock()

		time.Sleep(5 * utils.T_QUANT * time.Millisecond)
	}
}

// ----------------------------------------------------------------------------------------------------|

func renderNameInput(name []rune, canvas *rendering.Canvas) {
	if len(name) > 8 {
		name = name[:8]
	}

	tmpImg := rendering.Img([]string{"        "})
	canvas.SetImage(tmpImg, 10, 13)
	tmpImg = rendering.Img([]string{string(name)})
	canvas.SetImage(tmpImg, 10, 13)

	canvas.Print()
}
