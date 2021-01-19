package main

import (
	"fmt"
	"strings"
)

// Game holds the Cells of the game of life.
type Game struct {
	// Current is the current game state.
	Current [][]Cell
	// buffer for computing a new state.
	buffer [][]Cell
}

// Cell is alive when true or dead when false.
type Cell bool

// String returns "X" in case of a living Cell
// and " " in case of a dead Cell.
func (c Cell) String() string {
	if c {
		return "X"
	} else {
		return " "
	}
}

// Int converts a living Cell into a one and a dead one into a zero.
func (c Cell) Int() int {
	if c {
		return 1
	} else {
		return 0
	}
}

// LivingNeighbours counts the numbers of living neighbour Cells in the top, right,
// bottom, left and each corner.
func (g Game) LivingNeighbours(x, y int) int {
	livingNeighbours := 0
	for yDelta := -1; yDelta <= 1; yDelta++ {
		for xDelta := -1; xDelta <= 1; xDelta++ {
			if yDelta == 0 && xDelta == 0 {
				continue
			}
			livingNeighbours += g.safeGet(x+xDelta, y+yDelta).Int()
		}
	}
	return livingNeighbours
}

// Update's a Cell considering its number of living neighbours.
func (c Cell) Update(livingNeighbours int) Cell {
	if !c && livingNeighbours == 3 {
		// Gets born
		return true
	}
	if c && livingNeighbours < 2 {
		// Dead by loneliness
		return false
	}
	if c && (livingNeighbours == 2 || livingNeighbours == 3) {
		// Stay alive
		return true
	}
	if c && livingNeighbours > 3 {
		// Overpopulation
		return false
	}
	// Keep status
	return c
}

// safeGet returns the Cell at the specified position
// or a dead Cell if the position is out of bounds.
func (g Game) safeGet(x, y int) Cell {
	// Check if index is out of bounds
	// Return false as default
	if y < 0 || y >= len(g.Current) {
		return false
	}
	if x < 0 || x >= len(g.Current[y]) {
		return false
	}
	return g.Current[y][x]
}

// Update computes a new state of the game.
func (g *Game) Update() {
	// Update each cell
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[y]); x++ {
			// Write the updated cell to the buffer
			g.buffer[y][x] = g.Current[y][x].Update(g.LivingNeighbours(x, y))
		}
	}

	// Reassign the current game
	g.Current = g.buffer
}

// NewGame creates a new game.
func NewGame(game [][]Cell) {
	// Create buffer of same size as the game
	buffer := make([][]Cell, len(g))
	for i := range a {
		buffer[i] = make([]Cell, len(g[i]))
	}

	return Game {
		Current: game,
		buffer: buffer,
	}
}

// String returns a grid of the current game state.
func (g Game) String() string {
	output := make([]string, len(g.Current))
	for y := 0; y < len(g.Current); y++ {
		var line strings.Builder
		for x := 0; x < len(g.Current[y]); x++ {
			line.WriteString(g.Current[y][x].String())
		}
		output[y] = line.String()
	}
	return strings.Join(output, "\n")
}

func main() {
	g := NewGame{{false, false, false}, {false, false, false}, {true, true, true}}
	fmt.Println("Round 1")
	fmt.Println(g.String())
	g.Update()
	fmt.Println("Round 2")
	fmt.Println(g.String())
}
