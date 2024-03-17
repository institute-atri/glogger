# ✏️ GLOGGER - Logging and Colorizing Output

[![CI](https://github.com/institute-atri/glogger/actions/workflows/ci.yml/badge.svg)](https://github.com/institute-atri/glogger/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/institute-atri/glogger)](https://goreportcard.com/report/github.com/institute-atri/glogger)
[![codecov](https://codecov.io/gh/institute-atri/glogger/graph/badge.svg?token=2eXFFi17z2)](https://codecov.io/gh/institute-atri/glogger)

Simplifying the usage of loggers in files and standard output of the terminal/command prompt

# ✏️ Usage

Displaying colorized output on stdout:

```go
package main

import (
	"github.com/institute-atri/glogger"
)

func main() {
	glogger.Done("Hello DONE")
	glogger.Danger("Hello DANGER")
	glogger.Warning("Hello WARNING")
	glogger.Info("Hello INFO")
	glogger.Fatal("Hello FATAL")
}
```

Example illustrating the usage of glogger to create log files:

```go
package main

import (
	"github.com/institute-atri/glogger"
)

func main() {
	logger := glogger.NewLogger("MAIN FUNC", false)

	err := logger.Done("Hello DONE")
	glogger.ErrorHandling(err)
}
```

Describing the `NewLogger(prefix string, silent bool) *Logger` function:

1. The 'prefix' argument enables you to specify a string value identifying the location of the error or logger output.
2. The 'silent' parameter determines whether the same outputs appearing in the log file will also be displayed in the
   stdout of the terminal/command prompt.

The `NewLogger` function is responsible for writing a log file, providing enhanced details about the software process or
potential errors.</p>

We can also write directly to the log file using the `WriterLogsFile(s string) error` function.

```go
package main

import (
	"github.com/institute-atri/glogger"
)

func main() {
	logger := glogger.NewLogger("MAIN FUNC", false)

	err := logger.WriterLogsFile("Hello Logger")
	glogger.ErrorHandling(err)
}
```

# ✏️ License

This project is licensed under the **MIT License**—see the [LICENSE](LICENSE) file.

```text
MIT License

Copyright (c) 2024 ATRI - Advanced Technology Research Institute

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```