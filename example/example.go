// test program for the readline package
package main

import (
	"fmt"
	"strings"
	"github.com/gobs/readline"
)

var (
	words = []string { "alpha", "beta", "charlie", "delta", "another", "banana", "carrot", "delimiter" }
	matches = make([]string, 0, len(words))
)

func AttemptedCompletion(text string, start, end int) (result []string) {
	result = make([]string, 0, len(words))

	for _, w := range words {
		if strings.HasPrefix(w, text) {
			result = append(result, w)
		}
	}

	return
}

func CompletionEntry(prefix string, index int) string {
	if index == 0 {
		matches = matches[:0]

		for _, w := range words {
			if strings.HasPrefix(w, prefix) {
				matches = append(matches, w)
			}
		}
	}

	if index < len(matches) {
		return matches[index]
	} else {
		return ""
	}
}

func main() {
	prompt := "by your command> ";

	// this crashes at times (at least on Windows)
	//readline.SetAttemptedCompletionFunction(AttemptedCompletion)

	// this seems to work more reliably
	readline.SetCompletionEntryFunction(CompletionEntry)

	// loop until ReadLine returns nil (signalling EOF)
	L: for {
		result := readline.ReadLine(&prompt);
		if result == nil { // exit loop
			break L
		}

		line := *result

		switch line {
		case "": // ignore blank lines
			continue

		case "exit", "quit":
			break L

		default:
			if strings.HasPrefix(line, "prompt ") {
				prompt = strings.TrimPrefix(line, "prompt ")
			} else {
				fmt.Println(line);
			}

			readline.AddHistory(line); //allow user to recall this line
		}
	}
}
