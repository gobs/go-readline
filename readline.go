// Wrapper around the GNU readline(3) library

// +build linux darwin windows,libreadline

package readline

/*
 #cgo LDFLAGS: -lreadline

 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
 #include "readline/readline.h"
 #include "readline/history.h"

 #if !defined(RL_READLINE_VERSION) || (RL_READLINE_VERSION < 0x0600)
 typedef char **rl_completion_func_t (const char *, int, int);
 #endif

 extern void set_completion_entry_function();
 extern void set_attempted_completion_function();

 extern char **cstring_array_new(int);
 extern void cstring_array_set(char **, int, char*);

 extern char *null_cstring();
 extern char **null_cstring_array();
*/
import "C"
import "unsafe"
import "syscall"

// Read a line
func ReadLine(prompt *string) *string {
	var p *C.char

	//readline allows an empty prompt(NULL)
	if prompt != nil {
		p = C.CString(*prompt)
	}

	ret := C.readline(p)

	if p != nil {
		C.free(unsafe.Pointer(p))
	}

	if ret == nil {
		return nil
	} //EOF

	s := C.GoString(ret)
	C.free(unsafe.Pointer(ret))
	return &s
}

// Add line to history
func AddHistory(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.add_history(p)
}

// Parse and execute single line of a readline init file.
func ParseAndBind(s string) {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	C.rl_parse_and_bind(p)
}

// Parse a readline initialization file.
// The default filename is the last filename used.
func ReadInitFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.rl_read_init_file(p)
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

// Load a readline history file.
// The default filename is ~/.history.
func ReadHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.read_history(p)
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

var (
	HistoryLength = -1
)

// Save a readline history file.
// The default filename is ~/.history.
func WriteHistoryFile(s string) error {
	p := C.CString(s)
	defer C.free(unsafe.Pointer(p))
	errno := C.write_history(p)
	if errno == 0 && HistoryLength >= 0 {
		errno = C.history_truncate_file(p, C.int(HistoryLength))
	}
	if errno == 0 {
		return nil
	}
	return syscall.Errno(errno)
}

// Set the readline word delimiters for tab-completion
func SetCompleterDelims(break_chars string) {
	p := C.CString(break_chars)
	//defer C.free(unsafe.Pointer(p))
	C.free(unsafe.Pointer(C.rl_completer_word_break_characters))
	C.rl_completer_word_break_characters = p
}

// Get the readline word delimiters for tab-completion
func GetCompleterDelims() string {
	cstr := C.rl_completer_word_break_characters
	delims := C.GoString(cstr)
	return delims
}

//////////////////////////////////////////////////////////////////////////////////

// The signature for the rl_completion_entry_function callback
type go_compentry_func_t func(text string, state int) string

// The signature for the rl_attempted_completion_function callback
type go_completion_func_t func(text string, start, end int) []string

var _go_completion_entry_function go_compentry_func_t
var _go_attempted_completion_function go_completion_func_t

//export go_CompletionEntryFunction
func go_CompletionEntryFunction(text *C.char, state int) *C.char {
	if _go_completion_entry_function != nil {
		ret := _go_completion_entry_function(C.GoString(text), state)
		if len(ret) > 0 {
			return C.CString(ret)
		}
	}

	return C.null_cstring()
}

//export go_AttemptedCompletionFunction
func go_AttemptedCompletionFunction(text *C.char, start, end int) **C.char {
	if _go_attempted_completion_function != nil {
		ret := _go_attempted_completion_function(C.GoString(text), start, end)
		if ret != nil {
			size := len(ret)
			c_ret := C.cstring_array_new(C.int(size + 1))

			for i, s := range ret {
				C.cstring_array_set(c_ret, C.int(i), C.CString(s))
			}

			C.cstring_array_set(c_ret, C.int(size), C.null_cstring())
			return c_ret
		}
	}

	return C.null_cstring_array()
}

// Set rl_completion_entry_function
func SetCompletionEntryFunction(cbk go_compentry_func_t) {
	_go_completion_entry_function = cbk
	C.set_completion_entry_function()
}

// Set rl_attempted_completion_function
func SetAttemptedCompletionFunction(cbk go_completion_func_t) {
	_go_attempted_completion_function = cbk
	C.set_attempted_completion_function()
}

/* EOF */
