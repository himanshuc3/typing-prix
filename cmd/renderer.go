package cmd

import (
	"fmt"

	"github.com/gosuri/uilive"
)

var writer = uilive.New()

func Print(text string) {
	writer.Start()
	clear()
	fmt.Fprintln(writer, text)
}

func Stop() {
	writer.Stop()
}

func Update(text string) {
	fmt.Fprintln(writer, text)
}

func clear() {
	// NOTE:
	// 1. ANSI escape sequence for cursor location,
	// color, font-styling etc. Mimicking clearing buffer here
	fmt.Print("\033[H\033[2J")
}
