package timeformat_test

import (
	"fmt"
	"time"

	"github.com/v4run/timeformat"
)

func ExampleLayout() {
	fmt.Println(timeformat.Layout("www MMM D HH:mm:ss -ZZ z YYYY"))
	// Output: Mon Jan 2 15:04:05 -0700 MST 2006
}

func ExampleT() {
	t := time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(timeformat.T(t).Format("www MMM D HH:mm:ss -ZZ z YYYY"))
	// Output: Sun Jan 1 00:00:00 +0000 UTC 2017
}
