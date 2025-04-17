package write

import (
	"fmt"

	"github.com/gosuri/uilive"
)

type LiveRenderer struct {
	writer    *uilive.Writer
	initiated bool
}

// NOTE:
// 1. Variables declared globally are accessible on a package level
// rather than only to the file
// 2. uilive - Extraction over os.Stdout writer for flushing and synchronized writes
var Writer = New()

func New() *LiveRenderer {
	return &LiveRenderer{
		writer:    uilive.New(),
		initiated: false,
	}
}

func (renderer *LiveRenderer) Print(text string) {
	if renderer.initiated {
		fmt.Fprintf(renderer.writer, text+"\n")
		return
	}
	renderer.writer.Start()
	renderer.initiated = true
	fmt.Fprintf(renderer.writer, text+"\n")
}

func (renderer *LiveRenderer) Stop() {
	renderer.writer.Stop()
}

func clear() {
	// NOTE:
	// 1. ANSI escape sequence for cursor location,
	// color, font-styling etc. Mimicking clearing buffer here
	fmt.Print("\033[H\033[2J")
}
