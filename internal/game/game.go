package game

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
		
	}
	writer.Print(game.)
	go game.startTimer()
	return game
}

func (g *Game) get