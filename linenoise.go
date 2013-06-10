package linenoise

// -windows

// #include <stdlib.h>
// #include "linenoise.h"
// #include "linenoiseCompletionCallbackHook.h"
import "C"

import (
	"errors"
	"fmt"
	"unsafe"
)

func init() {
	C.linenoiseSetupCompletionCallbackHook()
}

// Line displays given string and returns line from user input.
func Line(prompt string) string { // char *linenoise(const char *prompt);
	promptCString := C.CString(prompt)
	resultCString := C.linenoise(promptCString)
	C.free(unsafe.Pointer(promptCString))

	result := C.GoString(resultCString)
	C.free(unsafe.Pointer(resultCString)) // TODO: is this required?
	return result
}

// AddHistory adds a line to history. Returns non-nil error on fail.
func AddHistory(line string) error { // int linenoiseHistoryAdd(const char *line);
	lineCString := C.CString(line)
	res := C.linenoiseHistoryAdd(lineCString)
	C.free(unsafe.Pointer(lineCString))
	if res != 1 {
		return errors.New("Could not add line to history.")
	}
	return nil
}

// SetHistoryCapacity changes the maximum length of history. Returns non-nil error on fail.
func SetHistoryCapacity(capacity int) error { // int linenoiseHistorySetMaxLen(int len);
	res := C.linenoiseHistorySetMaxLen(C.int(capacity))
	if res != 1 {
		return errors.New("Could not set history max len.")
	}
	return nil
}

// SaveHistory saves from file with given filename. Returns non-nil error on fail.
func SaveHistory(filename string) error { // int linenoiseHistorySave(char *filename);
	filenameCString := C.CString(filename)
	res := C.linenoiseHistorySave(filenameCString)
	C.free(unsafe.Pointer(filenameCString))
	if res != 0 {
		return errors.New("Could not save history to file.")
	}
	return nil
}

// LoadHistory loads from file with given filename. Returns non-nil error on fail.
func LoadHistory(filename string) error { // int linenoiseHistoryLoad(char *filename);
	filenameCString := C.CString(filename)
	res := C.linenoiseHistoryLoad(filenameCString)
	C.free(unsafe.Pointer(filenameCString))
	if res != 0 {
		return errors.New("Could not load history from file.")
	}
	return nil
}

// Clear clears the screen.
func Clear() { // void linenoiseClearScreen(void);
	C.linenoiseClearScreen()
}

// SetMultiline sets linenoise to multiline or single line.
// In multiline mode the user input will be wrapped to a new line when the length exceeds the amount of available rows in the terminal.
func SetMultiline(ml bool) { // void linenoiseSetMultiLine(int ml);
	if ml {
		C.linenoiseSetMultiLine(1)
	} else {
		C.linenoiseSetMultiLine(0)
	}
}

// Completions defines a list of completions for linenoise. It implements Completer
type Completions struct {
	// typedef struct linenoiseCompletions {
	//   size_t len;
	//   char **cvec;
	// } linenoiseCompletions;
	compl *C.linenoiseCompletions
}

//export linenoiseGoCompletionCallbackHook
func linenoiseGoCompletionCallbackHook(input *C.char, compl *C.linenoiseCompletions) {
	fmt.Println("hook was called")
	//++ call actual handler (interface, closure, global var)
}

// typedef void(linenoiseCompletionCallback)(const char *, linenoiseCompletions *);
// void linenoiseSetCompletionCallback(linenoiseCompletionCallback *);
// void linenoiseAddCompletion(linenoiseCompletions *, char *);
