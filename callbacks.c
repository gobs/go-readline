// +build linux darwin libreadline

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "readline/readline.h"
#include "readline/history.h"

#if !defined(RL_READLINE_VERSION) || (RL_READLINE_VERSION < 0x0600)
typedef char **rl_completion_func_t (const char *, int, int);
#endif

#include "_cgo_export.h"

void set_completion_entry_function() {
	rl_completion_entry_function = (rl_compentry_func_t *)go_CompletionEntryFunction;
}

void set_attempted_completion_function() {
	rl_attempted_completion_function = (rl_completion_func_t *)go_AttemptedCompletionFunction;
}

char *null_cstring() {
	return (char *)0;
}

char **null_cstring_array() {
	return (char **)0;
}

char **cstring_array_new(int size) {
	return (char **) malloc(size * sizeof(char *));
}

void cstring_array_set(char **csa, int i, char *s) {
	csa[i] = s;
}
