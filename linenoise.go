package linenoise

// +linux

// #include <stdlib.h>
// #include "linenoise.h"
import "C"

import (
	"unsafe"
)

// Line displays given string and returns line from user input.
// char *linenoise(const char *prompt);
func Line(prompt string) string {
	promptCString := C.CString(prompt)
	resultCString := C.linenoise(promptCString) // will error on this line
	C.free(promptCString)                       // but should error on this line

	result := C.GoString(resultCString)
	C.free(unsafe.Pointer(resultCString)) // TODO: is this required?
	return result
}

// int linenoiseHistoryAdd(const char *line);
// int linenoiseHistorySetMaxLen(int len);
// int linenoiseHistorySave(char *filename);
// int linenoiseHistoryLoad(char *filename);
// void linenoiseClearScreen(void);
// void linenoiseSetMultiLine(int ml);

// typedef struct linenoiseCompletions {
//   size_t len;
//   char **cvec;
// } linenoiseCompletions;

// typedef void(linenoiseCompletionCallback)(const char *, linenoiseCompletions *);
// void linenoiseSetCompletionCallback(linenoiseCompletionCallback *);
// void linenoiseAddCompletion(linenoiseCompletions *, char *);
