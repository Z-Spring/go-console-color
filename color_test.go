package go_console_color

import (
	"fmt"
	"testing"
)

func TestConsoleColor(t *testing.T) {
	s := ConsoleColor(Cyan, "hello world!")
	fmt.Println(s)
}
