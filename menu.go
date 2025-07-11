package gameparts

import (
	"strconv"
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"
	"time"

	"github.com/nsf/termbox-go"
)

// ----------------------------------------------------------------------------------------------------|

func Menu(control *utils.Controls, canvas *rendering.Canvas) int {
	canvas.Clear()

	newImg := rendering.CreateImgLogo()
	canvas.SetImage(newImg, 30, 8)
	newImg = rendering.CreateImgLevel()
	canvas.SetImage(newImg, 15, 20)

	canvas.Print()

	return getLevel(control, canvas)
}

// ----------------------------------------------------------------------------------------------------|

func getLevel(control *utils.Controls, canvas *rendering.Canvas) int {
	var levelStr string

	for {
		control.MutexEvent.Lock()
		if control.NewEvent {

			if control.Ev.Type == termbox.EventKey {
				switch {
				case control.Ev.Key == termbox.KeyEnter && levelStr != "":
					level, _ := strconv.Atoi(levelStr)
					if 0 <= level && level <= 9 {
						control.MutexEvent.Unlock()
						return level
					}

				case control.Ev.Key == termbox.KeyBackspace || control.Ev.Key == termbox.KeyBackspace2:
					levelStr = ""
					renderLevel(canvas, levelStr)

				case control.Ev.Ch >= '0' && control.Ev.Ch <= '9':
					levelStr = string(control.Ev.Ch)
					renderLevel(canvas, levelStr)
				}
			}

			control.NewEvent = false
		}
		control.MutexEvent.Unlock()

		time.Sleep(5 * utils.T_QUANT * time.Millisecond)
	}
}

// ----------------------------------------------------------------------------------------------------|

func renderLevel(canvas *rendering.Canvas, levelStr string) {
	canvas.SetImage([]string{" "}, 36, 20)

	if levelStr != "" {
		canvas.SetImage([]string{levelStr}, 36, 20)
	}

	canvas.Print()
}
