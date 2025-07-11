package gameparts

import (
	"io"
	"os"
	"sort"
	"terminal-tetris2/encryption"
	"terminal-tetris2/rendering"
	"terminal-tetris2/utils"

	"github.com/nsf/termbox-go"
)

// -----------------------------------------------------------------------------------------|

func EndGame(
	control *utils.Controls,
	canvas *rendering.Canvas,
	name string,
	level, score int,
) bool {

	allEntries, markIndex := getAllSortedEntries(name, level, score)

	saveResult(allEntries)
	topScores := getTopScores(utils.TOP_SCORES)

	// отрисовка доски с очками
	canvas.Clear()
	newImg := rendering.CreateImgBoard(topScores, markIndex)
	canvas.SetImage(newImg, 18, 3)
	newImg = rendering.CreateImgNextGame()
	canvas.SetImage(newImg, 5, 20)
	canvas.Print()

	return want2Exit(control)
}

// -----------------------------------------------------------------------------------------|

func saveResult(entries []utils.ScoreEntry) {

	rawBytes := encryption.GetBytes(entries)

	encryption.Hashing(rawBytes)

	file, err := os.OpenFile(utils.SCOREBOARD_FILE, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	file.Write(rawBytes)

	// for i := 0; i < len(rawBytes); i++ {
	// 	fmt.Fprintf(file, "%d\n", rawBytes[i])
	// }

}

// -----------------------------------------------------------------------------------------|

func getAllSortedEntries(name string, level, score int) ([]utils.ScoreEntry, int) {

	allEntries := readScores()
	newEntry := createNewEntry(name, level, score)
	allEntries = append(allEntries, newEntry)

	// Сортировка по убыванию очков
	sort.Slice(allEntries, func(i, j int) bool {
		return allEntries[i].Score.Value.(int) > allEntries[j].Score.Value.(int)
	})

	// Определение позиции новой записи
	markIndex := -1
	for i, entry := range allEntries {
		if entry.Name.Value == newEntry.Name.Value &&
			entry.Level.Value == newEntry.Level.Value &&
			entry.Score.Value == newEntry.Score.Value {
			markIndex = i
			break
		}
	}

	return allEntries, markIndex
}

// -----------------------------------------------------------------------------------------|

func createNewEntry(name string, level, score int) utils.ScoreEntry {
	res := utils.ScoreEntry{
		Score: utils.EntryParam{
			Type:  utils.TYPE_INT,
			Len:   4,
			Num:   1,
			Value: score,
		},
		Level: utils.EntryParam{
			Type:  utils.TYPE_INT,
			Len:   4,
			Num:   1,
			Value: level,
		},
		Name: utils.EntryParam{
			Type:  utils.TYPE_STRING,
			Len:   byte(len(name)),
			Num:   byte(len([]rune(name))),
			Value: name,
		},
	}

	return res
}

// -----------------------------------------------------------------------------------------|

func getTopScores(n int) []utils.ScoreEntry {
	entries := readScores()

	if n > len(entries) {
		n = len(entries)
	}
	return entries[:n]
}

// -----------------------------------------------------------------------------------------|

func readScores() []utils.ScoreEntry {
	data, err := readFile(utils.SCOREBOARD_FILE)
	if err != nil || data == nil || len(data) == 0 {
		return []utils.ScoreEntry{}
	}

	encryption.ReverseHashing(data)

	var entries []utils.ScoreEntry
	offset := 0

	for offset < len(data) {
		nameParam, nextOffset, ok := parseParam(data, offset, "string")
		if !ok {
			break
		}
		offset = nextOffset

		levelParam, nextOffset, ok := parseParam(data, offset, "int")
		if !ok {
			break
		}
		offset = nextOffset

		scoreParam, nextOffset, ok := parseParam(data, offset, "int")
		if !ok {
			break
		}
		offset = nextOffset

		entry := utils.ScoreEntry{
			Name:  nameParam,
			Level: levelParam,
			Score: scoreParam,
		}
		entries = append(entries, entry)
	}

	return entries
}

// -----------------------------------------------------------------------------------------|

func readFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// -----------------------------------------------------------------------------------------|

func parseParam(data []byte, offset int, expectedValueType string) (utils.EntryParam, int, bool) {
	if offset+3 > len(data) {
		return utils.EntryParam{}, offset, false
	}

	paramType := data[offset]
	paramLen := data[offset+1]
	paramNum := data[offset+2]

	valueStart := offset + 3
	valueEnd := valueStart + int(paramLen)
	if valueEnd > len(data) {
		return utils.EntryParam{}, offset, false
	}

	valueBytes := make([]byte, paramLen)
	copy(valueBytes, data[valueStart:valueEnd])

	value := encryption.Bytes2Interface(valueBytes, expectedValueType)
	if value == nil {
		return utils.EntryParam{}, offset, false
	}

	entryParam := utils.EntryParam{
		Type:  paramType,
		Len:   paramLen,
		Num:   paramNum,
		Value: value,
	}

	return entryParam, valueEnd, true
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
