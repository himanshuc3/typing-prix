package cmd

import (
	"os"

	"github.com/eiannone/keyboard"
	"github.com/himanshuc3/typing-prix/internal/game"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "lights-out",
	Short: "Start the game to test the typing speed.",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	gameInstance := game.New("Why am I attempting this project?")
	getKeyboardInput(gameInstance)
}

func getKeyboardInput(gameInstance *game.Game) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		keyboard.Close()
	}()

	// fmt.Println("Press ESC to retire")

	for {
		runeChar, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		// NOTE:
		// 1. Printing unicode representation and the key in laymen syntax
		// fmt.Printf("You pressed: rune %q, key %X\r\n", runeChar, key)
		if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC {
			break
		}

		isFinished := gameInstance.Input(runeChar)

		if isFinished {
			gameInstance.Stop()
			os.Exit(0)
			break
		}

	}
}
