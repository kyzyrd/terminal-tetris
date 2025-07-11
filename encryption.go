package encryption

import (
	"os"
	"terminal-tetris2/utils"
)

// -----------------------------------------------------------------------------------------|

func GetBytes(entries []utils.ScoreEntry) []byte {
	// для ускорения чтобы небыло перевыделения памяти при append
	numBytes := 0
	for _, entry := range entries {
		numBytes += int(entry.Name.Len)
		numBytes += 8
	}
	entryBytes := make([]byte, 0, numBytes)

	// start of encrypt, create slice of byte of entries
	for _, entry := range entries {
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

func Hashing(rawBytes []byte) {
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

			if a == b && len(rawBytes) > 1 {
				b = (b + 1) % len(rawBytes)
			}

			if len(rawBytes) > 1 {
				rawBytes[a], rawBytes[b] = rawBytes[b], rawBytes[a]
			}

			brushBits(&rawBytes[a], &rawBytes[b])

			bit1 := j % 8
			bit2 := (j + offset) % 8

			if bit1 == bit2 {
				bit2 = (bit2 + 1) % 8
			}
			swapBits(&rawBytes[a], &rawBytes[b], bit1, bit2)
		}
	}
}

// -----------------------------------------------------------------------------------------|

func ReverseHashing(rawBytes []byte) {
	key := ""
	if len(os.Args) > 1 {
		key = os.Args[1]
	}

	offset := 0
	for i := len(key) - 1; i >= 0; i-- {
		offset += int(key[i])
		for j := len(rawBytes) - 1; j >= 0; j-- {
			a := j
			b := (j + offset) % len(rawBytes)

			if a == b && len(rawBytes) > 1 {
				b = (b + 1) % len(rawBytes)
			}

			bit1 := j % 8
			bit2 := (j + offset) % 8

			if bit1 == bit2 {
				bit2 = (bit2 + 1) % 8
			}

			swapBits(&rawBytes[a], &rawBytes[b], bit1, bit2)

			brushBits(&rawBytes[a], &rawBytes[b])

			if len(rawBytes) > 1 {
				rawBytes[a], rawBytes[b] = rawBytes[b], rawBytes[a]
			}
		}
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

func Bytes2Interface(bytes []byte, expectedType string) any {
	switch expectedType {
	case "string":
		return string(bytes)
	case "int":
		if len(bytes) != 4 {
			return nil
		}

		var result int
		for i := len(bytes) - 1; i > -1; i-- {
			result <<= 8
			result |= int(bytes[i])
		}
		return result
	}

	return nil
}
