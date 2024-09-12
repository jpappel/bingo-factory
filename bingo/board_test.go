package bingo_test

import (
	"iter"
	"testing"

	"github.com/jpappel/bingo-factory/bingo"
)

func TestRows(t *testing.T) {
	g := bingo.Game{}

	testGame := func(size int, length int) {

		g.Checked = make([]bool, size)

		testGroup := func(name string, iter iter.Seq[[]bool]) {
			for i := range size {
				g.Checked[i] = false
			}
			for group := range iter {
				if len(group) != length {
					t.Logf("Mismatching %s length: %d != %d", name, length, len(group))
					t.FailNow()
				}

				for i := range length {
					if group[i] != false {
						t.Errorf("Incorrect value in %s!\n", name)
					}
				}
			}
			for i := range size {
				g.Checked[i] = true
			}
			for group := range iter {
				if len(group) != length {
					t.Logf("Mismatching %s length: %d != %d\n", name, length, len(group))
				}

				for i := range length {
					if group[i] != true {
						t.Errorf("Incorrect value in %s!\n", name)
					}
				}
			}
		}

		testGroup("row", g.Rows(length))
		testGroup("col", g.Cols(length))
		testGroup("diag", g.Diags(length))

	}

	t.Log("Testing Square Games")
	testGame(9, 3)
	testGame(25, 5)
	t.Log("Testing Non-Square Games")
	testGame(22, 2)
}
