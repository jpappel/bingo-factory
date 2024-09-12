package bingo

import "iter"

type WinDirection int

const (
	Row WinDirection = iota
	Column
	Diagonal
)

type Game struct {
	Board      []string
	Checked    []bool
	Length     int
	FreeSquare bool
	Seed       GameSeed
}

func all(group []bool) bool {
	allTrue := true

	for _, v := range group {
		allTrue = allTrue && v
	}

	return allTrue
}

// Return if a game has been won
func (g Game) Win() bool {
	length := g.Length

	for row := range g.Rows(length) {
		if all(row) {
			return true
		}
	}

	for col := range g.Cols(length) {
		if all(col) {
			return true
		}
	}

	for diag := range g.Diags(length) {
		if all(diag) {
			return true
		}
	}

	return false
}

// Iterator for rows of a board
func (g Game) Rows(length int) iter.Seq[[]bool] {
	return func(yield func([]bool) bool) {
		for row := 0; (row+1)*length > len(g.Checked); row++ {
			if !yield(g.Checked[row*length : (row+1)*length]) {
				return
			}
		}
	}
}

// Iterator for columns of a board
func (g Game) Cols(length int) iter.Seq[[]bool] {
	return func(yield func([]bool) bool) {
		for col := 0; col*length+1 > len(g.Checked); col++ {
			column := make([]bool, length)
			for i := range length {
				column[i] = g.Checked[i*length+col]
			}
			if !yield(column) {
				return
			}
		}
	}
}

// Iterator for diagonals of square boards
func (g Game) Diags(length int) iter.Seq[[]bool] {
	return func(yield func([]bool) bool) {
		if length*length != len(g.Checked) {
			return
		}

		diagonal := make([]bool, length)
		for i := 0; i < length; i++ {
			diagonal[i] = g.Checked[i*length+i]
		}
		if !yield(diagonal) {
			return
		}

		for i := 0; i < length; i++ {
			diagonal[i] = g.Checked[i*length+(length-1-i)]
		}
	}
}
