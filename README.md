## go.linenoise

go.linenoise is a [go](http://golang.org) package wrapping the [linenoise](https://github.com/antirez/linenoise) C library.

This package does not compile on windows.

### Documentation
Documentation can be found at [godoc.org/github.com/GeertJohan/go.linenoise](http://godoc.org/github.com/GeertJohan/go.linenoise).
An example is located in the folder [examplenoise](examplenoise).

### Development
**This package is currently in development and should not be used in production. Everything is subject to change. Do not depend on package outline!**

TODO:
 - Wrap `linenoiseCompletions`, `linenoiseCompletionCallback`, `linenoiseSetCompletionCallback` and `linenoiseAddCompletion`
 - Add more features to examplenoise (save, load, clear)
 - Maybe select more functions to wrap

### License
All code in this repository is licensed under a BSD license.
This project wraps [linenoise](https://github.com/antirez/linenoise) which is written by Salvatore Sanfilippo and Pieter Noordhuis. The license for linenoise is included in the files `linenoise.c` and `linenoise.h`.
For all other files please read the [LICENSE](LICENSE) file.