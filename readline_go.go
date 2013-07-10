// This is a basic implementation that doesn't rely on libreadline.
// Used for Windows, where you probably don't have the library, but can be disabled via 'libreadline' tag:
//
//   build -tags libreadline

// +build !linux,!darwin,!libreadline

package readline

import (
	"bufio"
	"io"
	"os"
)

var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

// Read a line
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

// Add line to history
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

// Get the current readline buffer
func GetLineBuffer() string {
	return ""
}

// The signature for the rl_completion_entry_function callback
type go_compentry_func_t func(text string, state int) string

// The signature for the rl_attempted_completion_function callback
type go_completion_func_t func(text string, start, end int) []string

// Call rl_completion_matches with the Go (compentry_function) callback
func CompletionMatches(text string, cbk go_compentry_func_t) []string {
	return nil
}

// Set rl_completion_entry_function
func SetCompletionEntryFunction(cbk go_compentry_func_t) {
}

// Set rl_attempted_completion_function
func SetAttemptedCompletionFunction(cbk func(text string, start, end int) []string) {
}

/* EOF */
