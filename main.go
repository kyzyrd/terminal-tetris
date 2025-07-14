package main

import (
	"io"
	"log"
	"os"
	"sync"
	"terminal-tetris2/encryption"
	"terminal-tetris2/gameparts"
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"

	"github.com/nsf/termbox-go"
)

// ----------------------------------------------------------------------------------------------------|

func main() {
	var newScoreEntry *utils.ScoreEntry
	var level int
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

	allEntries := readScores()
	for !quit {
		level = gameparts.Menu(&control, canvas)
		newScoreEntry = gameparts.Game(&control, canvas, level)
		allEntries, quit = gameparts.EndGame(&control, canvas, allEntries, newScoreEntry)
	}
	saveResult(allEntries)
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

// ----------------------------------------------------------------------------------------------------|

func readScores() []*utils.ScoreEntry {
	data, err := readFile(utils.SCOREBOARD_FILE)
	if err != nil || data == nil || len(data) == 0 {
		return make([]*utils.ScoreEntry, 0)
	}

	encryption.Unhashing(data)

	var allEntries []*utils.ScoreEntry
	corrupted := false
	if allEntries, corrupted = encryption.ConvertBytes2Entries(data); corrupted {
		log.Fatal("scoreboard is corrupted")
		return make([]*utils.ScoreEntry, 0)
	}

	return allEntries
}

// ----------------------------------------------------------------------------------------------------|

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ----------------------------------------------------------------------------------------------------|

func saveResult(allEntries []*utils.ScoreEntry) {
	rawBytes := encryption.ConvertEntries2Bytes(allEntries)

	encryption.Hashing(rawBytes)

	file, err := os.OpenFile(utils.SCOREBOARD_FILE, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	file.Write(rawBytes)
}
