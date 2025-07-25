<div align="center">
  <img src="img.png" alt="Beautiful header image">
</div> <br><br>

# Terminal Tetris

Terminal-based Tetris game written in Go, created as a learning exercise under mentorship of **[Lignigno](https://github.com/lignigno)**. It recreates the gameplay and aesthetics of the original Electronica 60 version, while taking advantage of Go's concurrency, strong typing, and modular design.

---

## Features

* Classic Tetris mechanics: 7 tetrominoes, line clearing, level progression
* Smooth rendering with `termbox-go` and a 2D canvas abstraction
* Game state management using typed structs and clean package layout
* Encrypted score saving using custom bit-swapping and byte manipulation
* Optional next-shape preview and control hints
* Adjustable starting level (0–9)

---

## Controls

| Key               | Action                     |
| ----------------- | -------------------------- |
| ← / →             | Move shape left/right      |
| ↑                 | Rotate shape clockwise     |
| ↓                 | Accelerate fall            |
| Space             | Instant drop               |
| n                 | Toggle next-shape preview  |
| h                 | Toggle control hints       |
| u                 | Increase level             |
| Esc               | Quit                       |
| 0–9               | Level select / Name input  |
| Enter / Backspace | Confirm / Delete character |
| Y/N or Д/Н        | Replay or Exit at endgame  |

---

## Project Structure

```
terminal-tetris2/
├── main.go               # Entry point
├── encryption/           # Score encoding/decoding
├── gameparts/            # Menu, gameplay, endgame logic
├── mechs/                # Gameplay mechanics
├── rendering/            # ASCII rendering and UI
├── shape/                # Tetrominoes and brick logic
├── utils/                # Constants, types, helpers
```

---

## Highlights

* **Modular design**: Each package has a focused responsibility
* **Typed structs**: `Canvas`, `Shape`, `ScoreEntry`, `GameVar`, etc.
* **Efficient rendering**: Canvas uses minimal redraws
* **Encrypted scoreboard**: Byte-level reversible hashing with key support
* **Game loop**: Timed with `UnixMilli()` for consistent updates
* **Configurable themes**: Green and B/W ASCII modes
* **Sorted leaderboard**: Top 15 scores stored securely

---

## Installation

```bash
# Prerequisite: Go 1.16 or higher

# Clone the repo
$ git clone https://github.com/kyzyrd/terminal-tetris.git
$ cd terminal-tetris

# Install dependencies
$ go get github.com/nsf/termbox-go

# Build the game
$ go build -o terminal-tetris

# Run (optional encryption key)
$ go run . "your-key"
```

---

## Score Storage & Encryption

Scores are stored in a local `scoreboard` file, encrypted using a reversible bit-swap algorithm. The encryption key is optional and taken from CLI arguments.

Each score entry contains:

* Player name (up to 8 ASCII chars)
* Level reached
* Final score

---

## Inspiration

This project is a tribute to the original 1984 **[Tetris](https://www.youtube.com/watch?v=O0gAgQQHFcQ&t=1s)** by **[Alexey Pajitnov](https://en.m.wikipedia.org/wiki/Alexey_Pajitnov)** on the **[Electronica 60](https://en.m.wikipedia.org/wiki/Elektronika_60)**. Despite its limitations, that version set the standard for minimalist, addictive gameplay. This terminal version reflects that same design philosophy.

---

## For Tomorrow School Download

``` bash
rm -rfv terminal-tetris && \
git clone https://github.com/kyzyrd/terminal-tetris.git && \
cd terminal-tetris && \
go build -o /Users/Shared/tetris main.go && \
rm -rfv ../terminal-tetris
```