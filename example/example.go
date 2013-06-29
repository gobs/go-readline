// test program for the readline package
package main

import (
	"fmt"
	"strings"
	"github.com/gobs/readline"
)

func main() {
	prompt := "by your command> ";

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
