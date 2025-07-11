package main

import (
	"sync"
	"terminal-tetris2/gameparts"
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"

	"github.com/nsf/termbox-go"
)

// ----------------------------------------------------------------------------------------------------|

func main() {
	var name string
	var score, level int
	var quit bool
	var control utils.Controls

	termbox.Init()
	defer termbox.Close()

	go CatchEvent(
		&control.NewEvent,
		&control.Ev,
		&control.MutexEvent,
	)

	canvas := rendering.GetNewCanvas(utils.WINDOW_WIDTH, utils.WINDOW_HEIGHT, rendering.DEFAULT)

	for !quit {
		level = gameparts.Menu(&control, canvas)
		name, level, score = gameparts.Game(&control, canvas, level)
		quit = gameparts.EndGame(&control, canvas, name, level, score)
	}
}

// ----------------------------------------------------------------------------------------------------|

// функция для отлова событий с клавиатуры
func CatchEvent(eventExist *bool, ev *termbox.Event, mutexEvent *sync.Mutex) {
	for {
		tmpEv := termbox.PollEvent()

		mutexEvent.Lock()
		*ev = tmpEv
		*eventExist = true
		mutexEvent.Unlock()
	}
}
