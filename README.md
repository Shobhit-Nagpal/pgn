# pgn

A Go module for parsing and working with chess games in PGN (Portable Game Notation) format.

## Installation

```bash
go get github.com/Shobhit-Nagpal/pgn
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/Shobhit-Nagpal/pgn"
)

func main() {
    // Parse a PGN string
    gameStr := `[Event "F/S Return Match"]
[Site "Belgrade, Serbia JUG"]
[Date "1992.11.04"]
[Round "29"]
[White "Fischer, Robert J."]
[Black "Spassky, Boris V."]
[Result "1/2-1/2"]

1. e4 e5 2. Nf3 Nc6 3. Bb5 1/2-1/2`

    game, err := pgn.New(gameStr)
    if err != nil {
        panic(err)
    }

    // Access game information
    fmt.Println("Event:", game.Event())
    fmt.Println("White Player:", game.White())
    fmt.Println("Black Player:", game.Black())
    fmt.Println("Result:", game.Result())
}
```

## Features

- Parse PGN format chess games
- Access standard PGN tags (Event, Site, Date, Round, White, Black, Result)
- Custom tag support
- Move tracking
- Game result handling

## API Reference

### Game Creation

- `New(pgn string) (*Game, error)`: Create a new game from PGN string

### Tag Operations

- `GetTag(name string) string`: Get the value of a specific tag
- `SetTag(tag, value string)`: Set a tag
- `TagPairs() map[string]string`: Get all tag pairs

### Standard Tag Accessors

- `Event() string`: Get the event name
- `Site() string`: Get the site/location
- `Round() string`: Get the round number/identifier
- `Date() string`: Get the game date
- `White() string`: Get the white player's name
- `Black() string`: Get the black player's name

### Game Result Methods

- `Result() string`: Get the game result
- `SetResult(result string)`: Set the game result
- `IsDraw() bool`: Check if the game ended in a draw
- `Winner() string`: Get the winner ("White", "Black", "Draw", or "Unknown")

### Move Management

- `GetMove(number int) *Move`: Get a specific move by number
- `SetMove(number int, move *Move)`: Set a move at a specific number
- `Moves() map[int]*Move`: Get all moves

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## Acknowledgements

This module was made by reading the PGN specification provided by [fsmosca](https://github.com/fsmosca/PGN-Standard).

## License

This module is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
