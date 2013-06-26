// Wrapper around the GNU readline(3) library
//
// Vanilla OSX doesn't have rl_completion_func_t
// +build !darwin

package readline

/*
 #cgo LDFLAGS: -lreadline

 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
 #include "readline/readline.h"
 #include "readline/history.h"
*/

import "C"
import "unsafe"
import "syscall"

//
func SetAttemptedCompletionFunction(cbk func(text string, start, end int) []string) {
	c_cbk := (*C.rl_completion_func_t)(unsafe.Pointer(&cbk))
	C.rl_attempted_completion_function = c_cbk
}
/* EOF */
