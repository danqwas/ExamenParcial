package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"flag"
)

var (
	n          = flag.Int("n", 8, "n queens, in an nxn grid")
	CantBeDone = errors.New("Can't be done")
)

type Board struct {
	grid [][]bool

	queens []int
}

func main() {
	flag.Parse()

	board := createBoard(*n)

	err := board.PlaceQueens(*n)
	fmt.Println(board.String())
	if err != nil {
		log.Fatalln("What the actual fuck:", err)
	}

}

// Simple, but used here
func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

// Returns an nxn grid off booleans
// The boolean is true if there's a queen placed there
func createBoard(n int) Board {
	grid := make([][]bool, n)
	for i, _ := range grid {
		row := make([]bool, n)
		grid[i] = row
	}

	locations := make([]int, n)
	for i := range locations {
		locations[i] = -1
	}

	return Board{
		grid:   grid,
		queens: locations,
	}
}

func (b *Board) PlaceQueens(n int) error {
	// check for freeness
	if n <= 0 {
		return nil
	}

	y := len(b.grid) - n

placements:
	for x := range b.grid[y] {

		if b.grid[y][x] {
			continue placements
		}

		for queenY, queenX := range b.queens {

			if queenX == -1 {
				break
			}

			if x == queenX || abs(x-queenX) == abs(y-queenY) {
				continue placements
			}
		}

		b.grid[y][x] = true
		b.queens[y] = x

		err := b.PlaceQueens(n - 1)

		if err != nil {

			b.grid[y][x] = false
			b.queens[y] = -1

			continue placements
		}

		return nil
	}

	return CantBeDone
}

func (b *Board) String() string {
	toPrint := make([]string, len(b.grid)+2)
	toPrint[0] = strings.Repeat("_", len(b.grid)+2)
	toPrint[len(toPrint)-1] = strings.Repeat("-", len(b.grid)+2)
	for y, row := range b.grid {
		rowText := make([]string, len(row)+2)
		rowText[0] = "|"
		rowText[len(rowText)-1] = "|"

		for x, element := range row {
			if element {
				rowText[x+1] = "Q"
			} else {
				rowText[x+1] = "x"
			}
		}

		toPrint[y+1] = strings.Join(rowText, "")
	}
	return strings.Join(toPrint, "\n")
}
