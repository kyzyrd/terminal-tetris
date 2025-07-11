# Terminal Tetris

## Project Overview

**Terminal Tetris** is a terminal-based implementation of the classic Tetris game, developed as an educational exercise to master the Go programming language, under the mentorship of **Lignigno**. This final iteration (approximately the 3th or 5th) leverages Go's concurrency, modularity, and type safety to deliver a robust and engaging game, using the `termbox-go` library for rendering and input handling. The project adheres to best practices to ensure maintainability, performance, and a nostalgic user experience. Below, we detail the good practices employed and pay homage to the original Tetris, created on the Electronica 60.

### Good Practices in the Project

- **Modular Architecture**: The codebase is organized into packages (`gameparts`, `mechs`, `utils`, `shape`, `rendering`, `encryption`), each with a clear purpose. For instance, `rendering` handles visual output (`canvas.go`, `createimgfield.go`), `mechs` manages gameplay mechanics (`linefuncs.go`, `applyaction.go`), and `encryption` secures score data, promoting separation of concerns and reusability.

- **Concurrency and Synchronization**: Go’s goroutines and mutexes (`sync.Mutex` in `main.go`, `utils/types.go`) enable thread-safe input handling (`CatchEvent`), ensuring responsive, non-blocking gameplay.

- **Type-Safe Data Structures**: Structs like `Canvas`, `GameVar`, `Shape`, and `ScoreEntry` (`rendering/canvas.go`, `utils/types.go`, `shape/shape.go`) provide strong typing and encapsulation. Methods such as `Canvas.SetImage`, `Shape.Move`, and `Shape.Rotate` offer clean interfaces, abstracting internal logic.

- **Centralized Constants**: Game parameters (`WINDOW_HEIGHT`, `BASE_SPEED`, `KEY_LEFT` in `utils/const.go`) and rendering themes (`DEFAULT`, `BLACK_AND_WHITE` in `canvas.go`) are centralized, facilitating easy configuration and consistent styling.

- **Efficient Rendering**: The `rendering` package optimizes terminal output:
  - `Canvas` (`canvas.go`) uses a 2D rune array for efficient screen updates, with methods like `Clear`, `Print`, and `SetImage` for precise control.
  - Functions like `CreateImgField` (`createimgfield.go`) and `CreateImgNS` (`createimgns.go`) minimize redraws using string manipulation and calculated bounds (`getBounds`).

- **Robust Error Handling**: File operations (`endgame.go`) and type assertions (`createimgboard.go`) include comprehensive error checks, handling edge cases like missing files or invalid data.

- **Optimized Algorithms**:
  - **Collision Detection**: `CheckCollision` (`utils/funcs.go`) uses a map for O(1) lookups to detect overlaps.
  - **Line Management**: `FindFullLines`, `ClearLines`, and `DropLines` (`mechs/linefuncs.go`) efficiently handle line clearing and dropping with sorted bricks (`SortBricksByY` in `shape/shape.go`).
  - **Image Creation**: `CreateImgNS` (`createimgns.go`) optimizes memory by calculating minimal image sizes.

- **Encapsulated Game Mechanics**: The `mechs` package (`applyaction.go`, `trymovedown.go`, `tryfreezing.go`) encapsulates movement, rotation, and freezing logic with collision checks, improving maintainability.

- **Custom Score Encryption**: The `encryption` package (`encryption.go`) implements a reversible bit-swapping algorithm (`Hashing`, `ReverseHashing`) for secure scoreboard storage, leveraging Go’s byte manipulation.

- **Time-Based Game Loop**: The game loop (`game.go`) uses `time.Now().UnixMilli()` with timers (`Fell`, `Step`, `Clear` in `utils/types.go`) for consistent pacing, adjustable via constants (`BASE_SPEED`, `FELL_TIME`).

- **User-Friendly Interface**: The `rendering` package enhances usability:
  - `CreateImgHint` (`createimghint.go`) displays clear control instructions.
  - `CreateImgScore` (`createimgscore.go`) formats scores with thousands marked by stars.
  - `CreateImgBoard` (`createimgboard.go`) highlights the player’s score in the leaderboard.

- **Iterative Development**: Under Lignigno’s mentorship, the project evolved through multiple iterations (around 3th or 5th), refining code quality and functionality based on feedback.

- **Debugging**: Learned debugging techniques using fmt.Printf and fmt.Fprintf to output real-time error messages and logs to different terminal ttys, enabling effective issue tracking during development.

### Homage to the Original Tetris

The original **Tetris**, created by **Alexey Pajitnov** in 1984 on the **Electronica 60**, a Soviet computer with limited resources, set a standard for elegant game design. **Terminal Tetris**, guided by **Lignigno**, pays tribute by recreating its core mechanics—falling tetrominoes, line clearing, and score tracking—in a terminal environment using `termbox-go`, capturing the minimalist, retro aesthetic of the original.

The Electronica 60 version excelled despite hardware constraints, with simple controls and addictive gameplay. This project reflects that spirit through ASCII-based visuals (`[]` for bricks, `.` for empty spaces in `createimgfield.go`), intuitive controls (`mechs/applyaction.go`), and visual feedback like blinking lines (`mechs/linefuncs.go`). Modern enhancements—concurrent input, encrypted scores, and customizable themes—build on Pajitnov’s vision, showcasing Go’s capabilities while preserving the classic charm.

## Features

- Classic Tetris gameplay with seven tetrominoes (O, I, L, J, S, Z, T) and rotation support.
- Collision detection for shapes, old bricks, and game boundaries (10x20 grid).
- Score tracking, level progression (up to 9), and encrypted scoreboard storage (`scoreboard` file).
- Toggleable next-shape preview and control hints.
- Dynamic shape fall speed based on level, with manual acceleration (down arrow) and instant drop (space).
- Terminal-based rendering with green or black-and-white themes.
- Interactive menu for level selection (0-9), gameplay, and endgame with name input (up to 8 characters) and leaderboard display (top 15 scores).
- Line clearing with blinking effect and automatic brick dropping.
- Game over when a new shape cannot spawn without collision.
- Encrypted score storage using a key-dependent bit-swapping algorithm.

## Installation

1. **Install Go**: Ensure Go (version 1.16 or later) is installed. Follow instructions at [golang.org](https://golang.org/doc/install).
2. **Clone Repository**:
   ```bash
   git clone https://github.com/kyzyrd/terminal-tetris.git
   cd terminal-tetris
   ```
3. **Install Dependencies**:
   ```bash
   go get github.com/nsf/termbox-go
   ```
4. **Build Game**:
   ```bash
   go build -o terminal-tetris
   ```
5. **Run Game**:
   Optionally specify an encryption key:
   ```bash
   ./terminal-tetris "-"
   ```

## Usage

- Launch the game and select a level (0-9) in the menu.
- Control falling shapes (see [Controls](#controls)).
- Clear lines by filling rows; cleared lines blink before removal, and remaining bricks drop.
- Earn points for cleared lines, with bonuses for level (3 points) and no next-shape preview (5 points).
- Game ends when a new shape cannot spawn.
- Enter a username (up to 8 characters) to save your score.
- View the leaderboard (top 15 scores) and choose to replay (Y/Д) or exit (N/Н).
- Scores are encrypted and saved to `scoreboard`.

## Controls

- **Left Arrow**: Move shape left.
- **Right Arrow**: Move shape right.
- **Up Arrow**: Rotate shape clockwise.
- **Space**: Drop shape instantly.
- **Down Arrow**: Accelerate shape fall.
- **N**: Toggle next-shape preview.
- **H**: Toggle control hints.
- **U**: Increase level (max 9).
- **Esc**: Quit game.
- **Menu/Endgame Input**:
  - **0-9**: Select level or enter name characters.
  - **Enter**: Confirm level or name.
  - **Backspace**: Delete last name character.
  - **Y/Д or N/Н**: Replay or exit at endgame.

## Project Structure

- **main.go**: Entry point, initializes game loop, and coordinates menu, gameplay, and endgame phases.
- **utils/**:
  - **const.go**: Defines constants (e.g., `WINDOW_HEIGHT=24`, `BASE_SPEED=600ms`, `BASE_SCORE=19`).
  - **funcs.go**: Utilities like `CheckCollision` for collision detection.
  - **initialize.go**: Initializes game state (field, shapes).
  - **types.go**: Defines structs (`GameVar`, `Controls`, `ScoreEntry`).
- **shape/**:
  - **shape.go**: Manages tetrominoes, rotation (`Rotate`), movement (`Move`), and brick sorting (`SortBricksByY`).
- **rendering/**:
  - **canvas.go**: Handles terminal rendering with `Canvas`, themes, and cursor control.
  - **createimgboard.go**: Generates leaderboard and replay prompt images.
  - **createimgfield.go**: Renders game field with bricks and shapes.
  - **createimghint.go**: Displays control hints.
  - **createimgmenu.go**: Creates logo and level selection images.
  - **createimgns.go**: Renders next-shape preview.
  - **createimgscore.go**: Displays score, level, and cleared lines.
- **mechs/**:
  - **applyaction.go**: Processes user input with collision checks.
  - **linefuncs.go**: Handles line detection (`FindFullLines`), clearing (`ClearLines`), blinking (`BlinkLines`), and dropping (`DropLines`).
  - **tryfreezing.go**: Manages shape freezing, new shape generation, and stats (`countStats`).
  - **trymovedown.go**: Controls automatic shape descent based on level speed.
- **gameparts/**:
  - **menu.go**: Implements level selection menu.
  - **game.go**: Manages main game loop, rendering, and name input.
  - **endgame.go**: Displays leaderboard, saves scores, and handles replay.
- **encryption/**:
  - **encryption.go**: Encrypts/decrypts scores with bit-swapping.

## Dependencies

- [termbox-go](https://github.com/nsf/termbox-go): Terminal interface library.
- Go standard libraries: `sync`, `math`, `math/rand`, `fmt`, `strings`, `time`, `sort`, `strconv`, `io`, `os`.