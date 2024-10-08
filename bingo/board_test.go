package bingo_test

import (
	"iter"
	"testing"

	"github.com/jpappel/bingo-factory/bingo"
)

func TestRows(t *testing.T) {
	g := bingo.Game{}

	testGame := func(size int, length int) {

		g.Board = make([]bingo.Tile, size)

		testGroup := func(name string, iter iter.Seq[[]bingo.Tile]) {
			for i := range size {
				g.Board[i].Checked = false
			}
			for group := range iter {
				if len(group) != length {
					t.Logf("Mismatching %s length: %d != %d", name, length, len(group))
					t.FailNow()
				}

				for i := range length {
					if group[i].Checked != false {
						t.Errorf("Incorrect value in %s!\n", name)
					}
				}
			}
			for i := range size {
				g.Board[i].Checked = true
			}
			for group := range iter {
				if len(group) != length {
					t.Logf("Mismatching %s length: %d != %d\n", name, length, len(group))
				}

				for i := range length {
					if group[i].Checked != true {
						t.Errorf("Incorrect value in %s!\n", name)
					}
				}
			}
		}

		testGroup("row", g.Rows())
		testGroup("col", g.Cols())
		testGroup("diag", g.Diags())

	}

	t.Log("Testing Square Games")
	testGame(9, 3)
	testGame(25, 5)
	t.Log("Testing Non-Square Games")
	testGame(22, 2)
}
