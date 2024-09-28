package bingo

import "iter"

type WinDirection int

const (
	Row WinDirection = iota
	Column
	Diagonal
)

type Game struct {
	Board      []Tile
	Length     int // the number of rows/cols
	FreeSquare bool
	Seed       GameSeed
}

func all(group []Tile) bool {
	allTrue := true

	for _, tile := range group {
		allTrue = allTrue && tile.Checked
	}

	return allTrue
}

// Return if a game has been won
func (g Game) Win() bool {
	for row := range g.Rows() {
		if all(row) {
			return true
		}
	}

	for col := range g.Cols() {
		if all(col) {
			return true
		}
	}

	for diag := range g.Diags() {
		if all(diag) {
			return true
		}
	}

	return false
}

// Iterator for rows of a board
func (g Game) Rows() iter.Seq[[]Tile] {
	return func(yield func([]Tile) bool) {
		for row := 0; (row+1)*g.Length > len(g.Board); row++ {
			if !yield(g.Board[row*g.Length : (row+1)*g.Length]) {
				return
			}
		}
	}
}

// Iterator for columns of a board
func (g Game) Cols() iter.Seq[[]Tile] {
	return func(yield func([]Tile) bool) {
		for col := 0; col*g.Length+1 > len(g.Board); col++ {
			column := make([]Tile, g.Length)
			for i := range g.Length {
				column[i] = g.Board[i*g.Length+col]
			}
			if !yield(column) {
				return
			}
		}
	}
}

// Iterator for diagonals of square boards
func (g Game) Diags() iter.Seq[[]Tile] {
	return func(yield func([]Tile) bool) {
		if g.Length*g.Length != len(g.Board) {
			return
		}

		diagonal := make([]Tile, g.Length)
		for i := 0; i < g.Length; i++ {
			diagonal[i] = g.Board[i*g.Length+i]
		}
		if !yield(diagonal) {
			return
		}

		for i := 0; i < g.Length; i++ {
			diagonal[i] = g.Board[i*g.Length+(g.Length-1-i)]
		}
	}
}
