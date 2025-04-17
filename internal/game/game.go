package game

// NOTE:
// 1. Identifiers inside internal can be accessed by the whole
// module
type Game struct {
	inputText    string
	currentIndex int
	inputTextSlice []runes
	time         int
	words        int
	correctInput bool
}

func New(inputText string) *Game{
	game := &Game {
		inputText: inputText,
		currentIndex: 0,
		inputTextSlice: []rune(inputText),
		words: 1,
		correctInput: true,
	}
	writer.Print(game.inputText)
	go game.startTimer()
	return game
}

func (g *Game) get