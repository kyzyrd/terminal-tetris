package encryption

import (
	"os"
	"terminal-tetris2/utils"
)

// -----------------------------------------------------------------------------------------|

func ConvertEntries2Bytes(allEntries []*utils.ScoreEntry) []byte {
	// для ускорения чтобы небыло перевыделения памяти при append
	numBytes := 0
	for _, entry := range allEntries {
		numBytes += int(entry.Name.Len)
		numBytes += 8
	}
	entryBytes := make([]byte, 0, numBytes)

	// start of encrypt, create slice of byte of allEntries
	for _, entry := range allEntries {
		entryBytes = append(entryBytes, entry.Name.Type, entry.Name.Len, entry.Name.Num)
		entryBytes = append(entryBytes, Interface2Bytes(entry.Name.Value)...)
		entryBytes = append(entryBytes, entry.Level.Type, entry.Level.Len, entry.Level.Num)
		entryBytes = append(entryBytes, Interface2Bytes(entry.Level.Value)...)
		entryBytes = append(entryBytes, entry.Score.Type, entry.Score.Len, entry.Score.Num)
		entryBytes = append(entryBytes, Interface2Bytes(entry.Score.Value)...)
	}

	return entryBytes
}

// -----------------------------------------------------------------------------------------|
func ConvertBytes2Entries(rawBytes []byte) ([]*utils.ScoreEntry, bool) {
	allEntries := make([]*utils.ScoreEntry, 0, len(rawBytes)/12)

	var entry *utils.ScoreEntry
	t, l := utils.TYPE_NIL, byte(0)
	i := 0

	for i < len(rawBytes) {
		if t == utils.TYPE_NIL {
			if i+3 > len(rawBytes) {
				return nil, true
			}

			t, l = rawBytes[i], rawBytes[i+1]
			num := rawBytes[i+2]
			i += 3

			if t == utils.TYPE_STRING {
				entry = &utils.ScoreEntry{}
			}
			if entry == nil {
				return nil, true
			}

			switch t {
			case utils.TYPE_SCORE:
				entry.Score = utils.EntryParam{Type: t, Len: l, Num: num}
			case utils.TYPE_LEVEL:
				entry.Level = utils.EntryParam{Type: t, Len: l, Num: num}
			case utils.TYPE_STRING:
				entry.Name = utils.EntryParam{Type: t, Len: l, Num: num}
			default:
				return nil, true
			}

		} else {
			if i+int(l) > len(rawBytes) || entry == nil {
				return nil, true
			}

			value := Bytes2Interface(rawBytes, t, i, i+int(l))
			i += int(l)

			switch t {
			case utils.TYPE_SCORE:
				entry.Score.Value = value
				allEntries = append(allEntries, entry)
				entry = nil
			case utils.TYPE_LEVEL:
				entry.Level.Value = value
			case utils.TYPE_STRING:
				entry.Name.Value = value
			default:
				return nil, true
			}

			t = utils.TYPE_NIL
		}
	}

	if i != len(rawBytes) {
		return nil, true
	}

	return utils.Sort(allEntries), false
}

// -----------------------------------------------------------------------------------------|

func Hashing(rawBytes []byte) {
	if len(rawBytes) <= 1 {
		return
	}
	key := ""
	if len(os.Args) > 1 {
		key = os.Args[1]
	}

	offset := 0
	for i := 0; i < len(key); i++ {
		offset += int(key[i])
		for j := 0; j < len(rawBytes); j++ {
			a := j
			b := (j + offset) % len(rawBytes)
			processIndex(rawBytes, a, b, j, false)
		}
	}
}

// -----------------------------------------------------------------------------------------|

func Unhashing(rawBytes []byte) {
	if len(rawBytes) <= 1 {
		return
	}
	key := ""
	if len(os.Args) > 1 {
		key = os.Args[1]
	}

	offset := 0
	for i := 0; i < len(key); i++ {
		offset += int(key[i])
	}
	for i := len(key) - 1; i >= 0; i-- {
		for j := len(rawBytes) - 1; j >= 0; j-- {
			a := j
			b := (j + offset) % len(rawBytes)
			processIndex(rawBytes, a, b, j, true)
		}
		offset -= int(key[i])
	}
}

// -----------------------------------------------------------------------------------------|

func processIndex(raw []byte, a, b, j int, reverse bool) {
	if a == b && len(raw) > 1 {
		b = (b + 1) % len(raw)
	}

	bit1 := j % 8
	bit2 := (j + b) % 8
	if bit1 == bit2 {
		bit2 = (bit2 + 1) % 8
	}

	if reverse {
		swapBits(&raw[a], &raw[b], bit1, bit2)
		brushBits(&raw[b], &raw[a])
		if len(raw) > 1 {
			raw[a], raw[b] = raw[b], raw[a]
		}
	} else {
		if len(raw) > 1 {
			raw[a], raw[b] = raw[b], raw[a]
		}
		brushBits(&raw[a], &raw[b])
		swapBits(&raw[a], &raw[b], bit1, bit2)
	}
}

// -----------------------------------------------------------------------------------------|

func swapBits(num1, num2 *byte, bit1, bit2 int) {
	mask1 := 1 << bit1
	mask2 := 1 << bit2

	bitOnRight1 := (int(*num1) & mask1) >> bit1
	bitOnRight2 := (int(*num2) & mask2) >> bit2

	*num1 = byte((int(*num1) & (^mask1)) | (bitOnRight2 << bit1))
	*num2 = byte((int(*num2) & (^mask2)) | (bitOnRight1 << bit2))
}

// -----------------------------------------------------------------------------------------|

func brushBits(num1, num2 *byte) {
	mask := byte(0b10101010)
	*num1, *num2 = (*num1)&mask|(*num2)&(^mask), (*num2)&mask|(*num1)&(^mask)
}

// -----------------------------------------------------------------------------------------|

func Interface2Bytes(i any) []byte {
	switch v := i.(type) {
	case string:
		return []byte(v)
	case int:
		intBytes := make([]byte, 4)
		for i := 0; i < len(intBytes); i++ {
			intBytes[i] = byte(v & 0xFF)
			v >>= 8
		}
		return intBytes
	}

	return []byte{}
}

// -----------------------------------------------------------------------------------------|

func Bytes2Interface(bytes []byte, expectedType byte, from, to int) any {
	switch expectedType {
	case utils.TYPE_STRING:
		return string(bytes[from:to])
	case utils.TYPE_LEVEL, utils.TYPE_SCORE:
		var result int
		if from > to {
			for i := from; i < to; i++ {
				result |= (int(bytes[i]) & 0xFF) << ((to - i - 1) * 8)
			}
		} else {
			for i := to - 1; i > from-1; i-- {
				result <<= 8
				result |= int(bytes[i]) & 0xFF
			}
		}

		return result
	}

	return nil
}
