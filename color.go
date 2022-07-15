package go_console_color

import (
	"strconv"
)

// define output color
// or if you just want to display the text color
// just delete 97; || 90;
// example: green = "\033[42m"
// example: white  = "\033[47m"
const (
	green  = "\033[97;42m"
	white  = "\033[90;47m"
	yellow = "\033[90;43m"
	red    = "\033[97;41m"
	blue   = "\033[97;44m"
	cyan   = "\033[97;46m"
	purple = "\033[97;45m"
	gray   = "\033[97;47m"
	reset  = "\033[0m"
)

type color string

const (
	Green  color = "Green"
	Red    color = "Red"
	Cyan   color = "Cyan"
	Blue   color = "Blue"
	Yellow color = "Yellow"
	White  color = "White"
	Purple color = "Purple"
	Gray   color = "Gray"
)

type ConsoleColor struct {
	C           color
	CancelColor bool
}

func NewColor(c color) *ConsoleColor {
	return &ConsoleColor{
		C: c,
	}
}

func convert(cs *ConsoleColor, target interface{}) string {
	var (
		suffix string
		text   string
	)

	switch target := target.(type) {
	case int64:
		targetConv := strconv.Itoa(int(target))
		suffix = targetConv + reset
		text = targetConv
	case string:
		suffix = target + reset
		text = target
	case float64:
		targetConv := strconv.FormatFloat(target, 'E', -1, 64)
		suffix = targetConv + reset
		text = targetConv

	case bool:
		targetConv := strconv.FormatBool(target)
		suffix = targetConv + reset
		text = targetConv

	case byte:
		suffix = string(target) + reset
		text = string(target)

	}

	switch cs.C {
	case Red:
		return red + suffix
	case Green:
		return green + suffix
	case Cyan:
		return cyan + suffix
	case Blue:
		return blue + suffix
	case White:
		return white + suffix
	case Yellow:
		return yellow + suffix
	case Purple:
		return purple + suffix
	case Gray:
		return gray + suffix
	default:
		return reset + text
	}

}

// String outputs the target content with color
func (cs *ConsoleColor) String(target string) string {
	if cs.CancelColor {
		return reset + target
	} else {
		return convert(cs, target)

	}
}

func (cs *ConsoleColor) Int64(target int64) string {
	targetConv := strconv.Itoa(int(target))

	if cs.CancelColor {
		return reset + targetConv
	} else {
		return convert(cs, target)

	}
}

func (cs *ConsoleColor) Float64(target float64) string {
	targetConv := strconv.FormatFloat(target, 'E', -1, 64)

	if cs.CancelColor {
		return reset + targetConv
	} else {
		return convert(cs, target)

	}
}

func (cs *ConsoleColor) Bool(target bool) string {
	targetConv := strconv.FormatBool(target)

	if cs.CancelColor {
		return reset + targetConv
	} else {
		return convert(cs, target)

	}
}
