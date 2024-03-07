<h1 align="center">GLOGGER</h1>
<h4 align="center">A Golang package simplifying the usage of loggers in files and standard output of the terminal/command prompt.</h4>
<br>
<br>

<h3>Examples:</h3>
The usage is extremely straightforward, requiring only a few lines to start using glogger.

<br>

<p>Displaying colorized output on stdout:</p>

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
<br>
<p>Example illustrating the usage of glogger to create log files:</p>

```go
package main

import (
  "github.com/institute-atri/glogger"
)

func main() {
  var logger = glogger.NewLogger("MAIN FUNC", false)

  var err = logger.Done("Hello DONE")

  if err != nil {
    glogger.Fatal(err)
  }
}
```

Describing the <code>NewLogger(prefix string, silent bool) *logger</code> function:

1. The 'prefix' argument enables you to specify a string value identifying the location of the error or logger output.
2. The 'silent' parameter determines whether the same outputs appearing in the log file will also be displayed in the stdout of the terminal/command prompt.

<p>The <code>NewLogger</code> function is responsible for writing a log file, providing enhanced details about the software process or potential errors.</p>

<br>

<p>We can also write directly to the log file using the <code>WriterLogsFile(s string) error</code> function.</p>

```go
package main

import (
  "github.com/institute-atri/glogger"
)

func main() {
  var logger = glogger.NewLogger("MAIN FUNC", false)

  var err = logger.WriterLogsFile("Hello Logger")

  if err != nil {
    glogger.Fatal(err)
  }
}
```
<h5 align="center">ATRI@2024 | License <a href="https://github.com/institute-atri/glogger/blob/main/LICENSE">MIT</a>.</h5>