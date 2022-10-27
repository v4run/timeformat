package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/v4run/timeformat"
)

func main() {
	flag.Parse()
	fmt.Println(timeformat.Layout(strings.Join(flag.Args(), " ")))
}
