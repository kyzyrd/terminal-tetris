package gameparts

import (
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"

	"github.com/nsf/termbox-go"
)

// -----------------------------------------------------------------------------------------|

func EndGame(
	control *utils.Controls,
	canvas *rendering.Canvas,
	allEntries []*utils.ScoreEntry,
	newScoreEntry *utils.ScoreEntry,
) ([]*utils.ScoreEntry, bool) {
	var markIndex int
	allEntries, markIndex = utils.SortAppend(allEntries, newScoreEntry)

	// отрисовка доски с очками
	canvas.Clear()
	newImg := rendering.CreateImgBoard(allEntries, markIndex)
	canvas.SetImage(newImg, 18, 3)
	newImg = rendering.CreateImgNextGame()
	canvas.SetImage(newImg, 5, 20)
	canvas.Print()

	return allEntries, want2Exit(control)
}

// -----------------------------------------------------------------------------------------|

func want2Exit(control *utils.Controls) bool {
	for {
		control.MutexEvent.Lock()
		if control.NewEvent {

			if control.Ev.Type == termbox.EventKey {
				switch control.Ev.Ch {
				case 'y', 'Y', 'д', 'Д':
					control.MutexEvent.Unlock()
					return false
				case 'n', 'N', 'н', 'Н':
					control.MutexEvent.Unlock()
					return true
				}
			}

			control.NewEvent = false
		}
		control.MutexEvent.Unlock()
	}
}
