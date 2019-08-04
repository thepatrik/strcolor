package strcolor_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/thepatrik/strcolor"
)

func TestReadmeUsage(t *testing.T) {
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

	// coloring can be disabled...
	enabled := strcolor.Enabled()
	defer strcolor.SetEnabled(enabled)
	strcolor.SetEnabled(false)
	fmt.Println(strcolor.Magenta("I prefer living in color."))
}

func TestPrintNormalColors(t *testing.T) {
	fmt.Printf("Is %s color black?\n", strcolor.Black("this"))
	fmt.Printf("Is %s color red?\n", strcolor.Red("this"))
	fmt.Printf("Is %s color green?\n", strcolor.Green("this"))
	fmt.Printf("Is %s color yellow?\n", strcolor.Yellow("this"))
	fmt.Printf("Is %s color blue?\n", strcolor.Blue("this"))
	fmt.Printf("Is %s color magenta?\n", strcolor.Magenta("this"))
	fmt.Printf("Is %s color cyan?\n", strcolor.Cyan("this"))
	fmt.Printf("Is %s color white?\n", strcolor.White("this"))
}

func TestPrintBrightColors(t *testing.T) {
	fmt.Printf("Is %s color bright black?\n", strcolor.BrightBlack("this"))
	fmt.Printf("Is %s color bright red?\n", strcolor.BrightRed("this"))
	fmt.Printf("Is %s color bright green?\n", strcolor.BrightGreen("this"))
	fmt.Printf("Is %s color bright yellow?\n", strcolor.BrightYellow("this"))
	fmt.Printf("Is %s color bright blue?\n", strcolor.BrightBlue("this"))
	fmt.Printf("Is %s color bright magenta?\n", strcolor.BrightMagenta("this"))
	fmt.Printf("Is %s color bright cyan?\n", strcolor.BrightCyan("this"))
	fmt.Printf("Is %s color bright white?\n", strcolor.BrightWhite("this"))
}

func TestPrintDisabledColors(t *testing.T) {
	if !strcolor.Enabled() {
		t.Error("should be enabled")
	}
	enabled := strcolor.Enabled()
	defer strcolor.SetEnabled(enabled)
	strcolor.SetEnabled(false)
	if strcolor.Enabled() {
		t.Error("should be not be enabled")
	}
	fmt.Printf("Is %s color default?\n", strcolor.Black("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Red("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Green("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Yellow("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Blue("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Magenta("this"))
	fmt.Printf("Is %s color default?\n", strcolor.Cyan("this"))
	fmt.Printf("Is %s color default?\n", strcolor.White("this"))
}

func TestPrintPadding(t *testing.T) {
	s := fmt.Sprintf("Is this %10s correctly?\n", strcolor.Yellow("padded"))
	fmt.Print(s)
	if len(s) != 39 {
		t.Error("padding is not correct", s)
	}

	s = fmt.Sprintf("%-20s%s\n", strcolor.Blue("And"), strcolor.White("this?"))
	fmt.Print(s)
	if len(s) != 44 {
		t.Error("padding is not correct", s)
	}
}
