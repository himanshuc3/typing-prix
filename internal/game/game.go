package game

import (
	"fmt"

	"github.com/TwiN/go-color"
	"github.com/eiannone/keyboard"
	"github.com/himanshuc3/typing-prix/internal/write"
)

var quit = make(chan struct{})

// NOTE:
// 1. Identifiers inside internal can be accessed by the whole
// module
type Game struct {
	inputText       string
	currentIndex    int
	inputTextSlice  []rune
	outputTextSlice []rune
	time            int
	words           int
	correctInput    bool
	matchedOutput   []byte
}

func New(inputText string) *Game {
	game := &Game{
		inputText:       inputText,
		currentIndex:    0,
		inputTextSlice:  []rune(inputText),
		outputTextSlice: []rune{},
		matchedOutput:   []byte{},
		words:           1,
		correctInput:    true,
	}
	write.Writer.Print(game.inputText)
	// go game.startTimer()
	return game
}

// NOTE:
// 1. Cache number of complete words.
func (g *Game) getWPM() int {

	var correctWords int
	// var isWordMatch bool = true
	// for rChar, num := range g.outputTextSlice {

	// 	if unicode.IsSpace(g.inputTextSlice[num]) {

	// 	}
	// }close(

	return correctWords
}

func (g *Game) Input(letter rune) bool {
	fmt.Printf("%b", letter == rune(keyboard.KeySpace))
	if letter == rune(keyboard.KeyBackspace2) {
		g.currentIndex = max(g.currentIndex-1, 0)
	} else if letter != rune(keyboard.KeySpace) && g.inputTextSlice[g.currentIndex] == ' ' {
		// Don't do anything
		return false
	} else {
		g.currentIndex++
	}

	g.updateMatchedOutput(letter)
	write.Writer.Print(g.getOutputText())
	return g.currentIndex == len(g.inputTextSlice)
}

func (g *Game) updateMatchedOutput(letter rune) {
	if g.currentIndex == 0 {
		g.matchedOutput = []byte{}
	} else if letter == rune(keyboard.KeyBackspace) {
		g.matchedOutput = g.matchedOutput[:g.currentIndex]
	} else {
		var matchedByte byte = CHAR_MATCH_TYPE["UNMATCHED"]
		// If a separator
		switch {
		case letter == rune(keyboard.KeySpace):

			matchedByte = CHAR_MATCH_TYPE["SEPARATOR"]
		case letter == g.inputTextSlice[g.currentIndex-1]:
			matchedByte = CHAR_MATCH_TYPE["MATCHED"]
		}
		g.matchedOutput = append(g.matchedOutput, matchedByte)
	}
	// fmt.Printf("Matched output array %v", g.matchedOutput)
}

func (g *Game) colorizeText() string {

	cursor := color.With(color.Yellow, "|")

	var begin string

	for num, char := range g.matchedOutput {
		if char == CHAR_MATCH_TYPE["UNMATCHED"] {
			// fmt.Println("unmatched")
			begin += color.Reset + color.With(color.Red, string(g.inputTextSlice[num]))
		} else {
			begin += color.Reset + color.With(color.Green, string(g.inputTextSlice[num]))
		}
	}

	rest := color.With(color.Blue, g.inputText[g.currentIndex:len(g.inputText)])

	return begin + cursor + rest

}

func (g *Game) getOutputText() string {
	return g.colorizeText() + "WPM: 1"
}

func (g *Game) Stop() {
	write.Writer.Stop()
	// TODO:
	// 1. Stop Timer
	// 2. Save Stats
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
