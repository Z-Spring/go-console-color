package go_console_color

// define output color
// or if you  just want to  display the text color
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

const (
	Green  string = "Green"
	Red    string = "Red"
	Cyan   string = "Cyan"
	Blue   string = "Blue"
	Yellow string = "Yellow"
	White  string = "White"
	Purple string = "Purple"
	Gray   string = "Gray"
)

// ConsoleColor display console's text  background color
func ConsoleColor(inputColor string, target string) string {
	suffix := target + reset

	switch inputColor {
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
		return reset + target
	}
}
