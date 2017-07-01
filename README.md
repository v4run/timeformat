# timeformat [![GoDoc](https://godoc.org/github.com/v4run/timeformat?status.svg)](https://godoc.org/github.com/v4run/timeformat) [![Build Status](https://travis-ci.org/v4run/timeformat.svg?branch=master)](https://travis-ci.org/v4run/timeformat)

An easy to remember time format layout for go.

## Layout
|timeformat|default|
----|----
YYYY|2006
YY|06
MMMM|January
MMM|Jan
MM|01
M|1
DD|02
D|2
hh|03
HH|15
h|3
wwww|Monday
www|Mon
mm|04
m|4
ss|05
s|5
f|0 *(.fff = .000)*
F|9 *(.FFF = .999)*
a|pm
A|PM
z|MST
-Z:Z:Z|-07:00:00
Z:Z:Z|Z07:00:00
-Z:Z|-07:00
Z:Z|Z07:00
-ZZZ|-070000
ZZZ|Z070000
-ZZ|-0700
ZZ|Z0700
-Z|-07
Z|Z07

## Example 1

```go
package main

import (
	"fmt"
	"time"

	"github.com/v4run/timeformat"
)

func main() {
	t := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(timeformat.T(t).Format("www MMM D HH:mm:ss -ZZ z YYYY"))
}
```

#### Output
```Sun Jan 1 00:00:00 +0000 UTC 2017```

## Example 2

```go
package main

import (
	"fmt"

	"github.com/v4run/timeformat"
)

func main() {
	fmt.Println(timeformat.Layout("www MMM D HH:mm:ss -ZZ z YYYY"))
}
```

#### Output
```Mon Jan 2 15:04:05 -0700 MST 2006```
