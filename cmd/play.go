package cmd

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
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
	gameInstance := game.NewGame()
	getKeyboardInput(runningGame)
}

func getKeyboardInput() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer func() {
		keyboard.Close()
	}()

	fmt.Println("Press ESC to retire")

	for {
		character, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// NOTE:
		// 1. Printing unicode representation and the key in laymen syntax
		fmt.Printf("You pressed: rune %q, key %X\r\n", character, key)
		if key == keyboard.KeyEsc {
			break
		}

		isFinished := game.Stop()

		if isFinished {
			game.Stop()
			os.Exit(0)
			break
		}

	}
}
