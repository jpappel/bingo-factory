package bingo

type GameSeed string

type Generator interface {
	New(int, int) *Game
	SetSeed(int64)
	Seed() int64
}

type RandomGenerator struct {
	tiles  TilePool
	picker TilePicker
	seed   int64
}

func (g RandomGenerator) New(size int, length int) *Game {
	g.picker.Reset()

	board := make([]string, 0, size)
	checked := make([]bool, size)

	for _, tile := range g.picker.Iter(size) {
		board = append(board, tile)
	}

	game := new(Game)
	game.Board = board
	game.Checked = checked
	game.Length = length

	return game
}

func (g RandomGenerator) Seed() int64 {
	return g.seed
}

func (g RandomGenerator) SetSeed(seed int64) {
	g.seed = seed
}
