package strcolor

import "fmt"

var (
	// normal colors
	black   = []byte{'\033', '[', '3', '0', 'm'}
	red     = []byte{'\033', '[', '3', '1', 'm'}
	green   = []byte{'\033', '[', '3', '2', 'm'}
	yellow  = []byte{'\033', '[', '3', '3', 'm'}
	blue    = []byte{'\033', '[', '3', '4', 'm'}
	magenta = []byte{'\033', '[', '3', '5', 'm'}
	cyan    = []byte{'\033', '[', '3', '6', 'm'}
	white   = []byte{'\033', '[', '3', '7', 'm'}

	// bright colors
	brightblack   = []byte{'\033', '[', '3', '0', ';', '1', 'm'}
	brightred     = []byte{'\033', '[', '3', '1', ';', '1', 'm'}
	brightgreen   = []byte{'\033', '[', '3', '2', ';', '1', 'm'}
	brightyellow  = []byte{'\033', '[', '3', '3', ';', '1', 'm'}
	brightblue    = []byte{'\033', '[', '3', '4', ';', '1', 'm'}
	brightmagenta = []byte{'\033', '[', '3', '5', ';', '1', 'm'}
	brightcyan    = []byte{'\033', '[', '3', '6', ';', '1', 'm'}
	brightwhite   = []byte{'\033', '[', '3', '7', ';', '1', 'm'}
)

var (
	esc   = "\033["
	clear = esc + "0m"
)

var enabled = true

// SetEnabled controls if colors are enabled or not
func SetEnabled(b bool) {
	enabled = b
}

// Enabled returns if colors are enabled
func Enabled() bool {
	return enabled
}

// StrColor struct
type StrColor struct {
	Val   interface{}
	Color []byte
}

func (c StrColor) String() string {
	if enabled {
		return esc + string(c.Color) + fmt.Sprint(c.Val) + clear
	}
	return fmt.Sprint(c.Val)
}

// Black string color
func Black(s interface{}) StrColor {
	return StrColor{Val: s, Color: black}
}

// Red string color
func Red(s interface{}) StrColor {
	return StrColor{Val: s, Color: red}
}

// Green string color
func Green(s interface{}) StrColor {
	return StrColor{Val: s, Color: green}
}

// Yellow string color
func Yellow(s interface{}) StrColor {
	return StrColor{Val: s, Color: yellow}
}

// Blue string color
func Blue(s interface{}) StrColor {
	return StrColor{Val: s, Color: blue}
}

// Magenta string color
func Magenta(s interface{}) StrColor {
	return StrColor{Val: s, Color: magenta}
}

// Cyan string color
func Cyan(s interface{}) StrColor {
	return StrColor{Val: s, Color: cyan}
}

// White string color
func White(s interface{}) StrColor {
	return StrColor{Val: s, Color: white}
}

// BrightBlack string color
func BrightBlack(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightblack}
}

// BrightRed string color
func BrightRed(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightred}
}

// BrightGreen string color
func BrightGreen(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightgreen}
}

// BrightYellow string color
func BrightYellow(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightyellow}
}

// BrightBlue string color
func BrightBlue(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightblue}
}

// BrightMagenta string color
func BrightMagenta(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightmagenta}
}

// BrightCyan string color
func BrightCyan(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightcyan}
}

// BrightWhite string color
func BrightWhite(s interface{}) StrColor {
	return StrColor{Val: s, Color: brightwhite}
}
