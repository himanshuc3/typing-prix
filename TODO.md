## Problem Statement

Pull up messages, quotes and excerpts from the world of F1 and let user do a typing test over cli.

1. Fetch data from sources, show a timer or run a text till completion. Calculate wpm, which is calculated by correct words (not taking into account characters.)

### MVP

1. Fetch text excerpt.
2. Calculate writing speed.

### v1 version

1. Interactive cli.
2. More options/configuration.
3.

### Implementation

Display the text to be completed -> Listen for keystrokes -> add it to the userinput and color characters as per strict string match -> terminate when the time runs out or the text is completed.
(all while the buffer is being rerendered and flushed)
