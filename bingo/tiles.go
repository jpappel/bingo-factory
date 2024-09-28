package bingo

import (
	"iter"
	"math/rand"
	"slices"
)

type Tile struct {
	Text      string
	Checked   bool
	Checkable bool
}

type TilePool map[string][]string

func (pool TilePool) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, list := range pool {
			for _, tile := range list {
				if !yield(tile) {
					return
				}
			}
		}
	}
}

type TilePicker interface {
	All() iter.Seq[string]            // provides an iterator over an entire TilePool
	Iter(int) iter.Seq2[string, Tile] // provides an iterator over n elements of a TilePool
	Reset()                           // reset the internal state of a TilePicker
}

type RandomTilePicker struct {
	ChosenTags []string
	tilePool   TilePool
	rand       rand.Rand
}

func NewRandomTilePicker(tiles TilePool, r rand.Rand) *RandomTilePicker {
	tp := new(RandomTilePicker)
	tp.tilePool = tiles
	tp.rand = r

	return tp
}

// Iterate over all elements of a TilePool in a random order
func (tp RandomTilePicker) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		tiles := slices.Collect(tp.tilePool.All())
		if len(tiles) == 0 {
			return
		}

		tp.rand.Shuffle(len(tiles), func(i int, j int) {
			tiles[i], tiles[j] = tiles[j], tiles[i]
		})

		for _, tile := range tiles {
			if !yield(tile) {
				return
			}
		}
	}
}

// Iterator over a TilePool by choosing one tile per tag until pool is exhausted or size tiles have been yielded
func (tp RandomTilePicker) Iter(size int) iter.Seq2[string, Tile] {
	return func(yield func(string, Tile) bool) {
		if len(tp.tilePool) == 0 {
			return
		}

		tags := make([]string, len(tp.tilePool))
		for tag := range tp.tilePool {
			tags = append(tags, tag)
		}
		tp.rand.Shuffle(len(tags), func(i int, j int) {
			tags[i], tags[j] = tags[j], tags[i]
		})

		yielded := 0
		for _, tag := range tags {
			if yielded == size {
				return
			}

			list := tp.tilePool[tag]
			if len(list) == 0 {
				continue
			}

			text := list[tp.rand.Intn(len(list))]
			tile := Tile{Text: text}
			if !yield(tag, tile) {
				return
			}
			yielded++
		}

	}
}

func (tp RandomTilePicker) Reset() {
	tp.ChosenTags = make([]string, len(tp.ChosenTags))
}
