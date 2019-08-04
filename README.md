# strcolor
[![Build Status](https://travis-ci.org/thepatrik/strcolor.svg?branch=master)](https://travis-ci.org/thepatrik/strcolor) [![Go Report Card](https://goreportcard.com/badge/github.com/thepatrik/strcolor)](https://goreportcard.com/report/github.com/thepatrik/strcolor) [![GoDoc](https://godoc.org/github.com/thepatrik/strcolor?status.svg)](https://godoc.org/github.com/thepatrik/strcolor)

Golang library for printing with standard ANSI colors.

#### install
```
import "github.com/thepatrik/strcolor"
```

## Usage

Colorize fmt/log prints:
```
package main

import (
	"bytes"
	"fmt"

	"github.com/thepatrik/strcolor"
)

func main() {
	// colorize a print...
	fmt.Println("Hello ", strcolor.Magenta("World"))

	// with formatting...
	fmt.Printf("%s %s %s %s %s %s %s %s.\n",
		strcolor.Black("Life"),
		strcolor.Red("is"),
		strcolor.Green("like"),
		strcolor.Yellow("a"),
		strcolor.Blue("box"),
		strcolor.Magenta("of"),
		strcolor.Cyan(8), // and not only strings.
		strcolor.White("crayons"))

	// or with a buffer...
	var buffer bytes.Buffer
	quote := "\"Color is a power which directly influences the soul.\""
	author := "Wassily Kandinsky"
	buffer.WriteString(strcolor.Cyan(quote).String() + " /" + author)
	fmt.Println(buffer.String())

	// coloring can be programmatically enabled/disabled (enabled by default on TTY character devices)...
	strcolor.SetEnabled(false)
	defer strcolor.SetEnabled(true)
	fmt.Println(strcolor.Magenta("I prefer living in color."))
}
```
