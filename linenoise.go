package linenoise

// +linux

// #include <stdlib.h>
// #include "linenoise.h"
import "C"

import (
	"fmt"
)

// DoSomething does something
func DoSomething() {
	fmt.Println("hi")
}
