package linenoise

// +linux

// #include <stdlib.h>
// #include "linenoise.h"
import "C"

import (
	"errors"
	"unsafe"
)

// Line displays given string and returns line from user input.
// char *linenoise(const char *prompt);
func Line(prompt string) string {
	promptCString := C.CString(prompt)
	resultCString := C.linenoise(promptCString)
	C.free(unsafe.Pointer(promptCString))

	result := C.GoString(resultCString)
	C.free(unsafe.Pointer(resultCString)) // TODO: is this required?
	return result
}

// AddHistory adds a line to history
// int linenoiseHistoryAdd(const char *line);
func AddHistory(line string) error {
	lineCString := C.CString(line)
	res := C.linenoiseHistoryAdd(lineCString)
	C.free(unsafe.Pointer(lineCString))
	if res != 1 {
		return errors.New("Could not add line to history.")
	}
	return nil
}

// SetHistoryMaxLen changes the maximum length of history
// int linenoiseHistorySetMaxLen(int len);
func SetHistoryMaxLen(capacity int) error { // TODO: maybe rename to SetHistoryCapacity
	res := C.linenoiseHistorySetMaxLen(C.int(capacity))
	if res != 1 {
		return errors.New("Could not set history max len.")
	}
	return nil
}

// SaveHistory saves from file with given filename.
// int linenoiseHistorySave(char *filename);
func SaveHistory(filename string) error {
	filenameCString := C.CString(filename)
	res := C.linenoiseHistorySave(filenameCString)
	C.free(unsafe.Pointer(filenameCString))
	if res != 1 {
		return errors.New("Could not save history to file.")
	}
	return nil
}

// LoadHistory loads from file with given filename.
// int linenoiseHistoryLoad(char *filename);
func LoadHistory(filename string) error {
	filenameCString := C.CString(filename)
	res := C.linenoiseHistoryLoad(filenameCString)
	C.free(unsafe.Pointer(filenameCString))
	if res != 1 {
		return errors.New("Could not load history from file.")
	}
	return nil
}

// Clear clears the screen.
// void linenoiseClearScreen(void);
func Clear() {
	C.linenoiseClearScreen()
}

// SetMultiline sets linenoise to multiline or single line
// void linenoiseSetMultiLine(int ml);
func SetMultiline(ml bool) {
	if ml {
		C.linenoiseSetMultiLine(1)
	} else {
		C.linenoiseSetMultiLine(0)
	}
}

// typedef struct linenoiseCompletions {
//   size_t len;
//   char **cvec;
// } linenoiseCompletions;

// typedef void(linenoiseCompletionCallback)(const char *, linenoiseCompletions *);
// void linenoiseSetCompletionCallback(linenoiseCompletionCallback *);
// void linenoiseAddCompletion(linenoiseCompletions *, char *);
