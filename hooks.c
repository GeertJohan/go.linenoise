#include "hooks.h"

extern void linenoiseGoCompletionCallbackHook(const char *, linenoiseCompletions *);
extern char * linenoiseGoHintCallbackHook(const char *, int *, int *);

void linenoiseSetupCompletionCallbackHook() {
	linenoiseSetCompletionCallback(linenoiseGoCompletionCallbackHook);
}

void linenoiseSetupHintCallbackHook() {
  linenoiseSetHintsCallback(linenoiseGoHintCallbackHook);
}
