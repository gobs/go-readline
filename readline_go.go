// +build windows

package readline

import (
	"bufio"
	"io"
	"os"
	)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

func ReadLine(prompt *string) *string {
	//readline allows an empty prompt(NULL)
	if prompt != nil && len(*prompt) > 0 {
		io.WriteString(os.Stdout, *prompt)
	}

	if !scanner.Scan() {
		return nil
	}

	s := scanner.Text()
	if s == "\x04" { // ^D
		return nil
	}

	return &s
}

func AddHistory(s string) {
}

// Parse and execute single line of a readline init file.
func ParseAndBind(s string) {
}

// Parse a readline initialization file.
// The default filename is the last filename used.
func ReadInitFile(s string) error {
	return nil
}

// Load a readline history file.
// The default filename is ~/.history.
func ReadHistoryFile(s string) error {
	return nil
}

var (
	HistoryLength = -1
)

// Save a readline history file.
// The default filename is ~/.history.
func WriteHistoryFile(s string) error {
	return nil
}

// Set the readline word delimiters for tab-completion
func SetCompleterDelims(break_chars string) {
}

// Get the readline word delimiters for tab-completion
func GetCompleterDelims() string {
	return ""
}

// 
func CompletionMatches(text string, cbk func(text string, state int) string) []string {
	return []string{}
}

//
func SetAttemptedCompletionFunction(cbk func(text string, start, end int) []string) {
}
/* EOF */
