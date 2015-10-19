# strm-skip

A driver for the **go-strm** Go programming language library, that provides the **SKIP** Strmer command.

**SKIP** skips rows from the beginning of the data table stream, and returns everything after that.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/strm-skip

[![GoDoc](https://godoc.org/github.com/reiver/strm-skip?status.svg)](https://godoc.org/github.com/reiver/strm-skip)

## Example
```
package main

import (
	. "github.com/reiver/strm-csv"
	. "github.com/reiver/strm-skip"
	. "github.com/reiver/strm-stdout"
)

func main() {
	Begin(CSV, "table.csv").
		Strm(SKIP, 12).
	End(STDOUT, "tsv")
}
```

(Note that in that example dot imports were used.)

## See Also

For more information about **go-strm** and for a list of other drivers, see:
https://github.com/reiver/go-strm
