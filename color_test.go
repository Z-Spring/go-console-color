package go_console_color

import (
	"fmt"
	"testing"
)

func TestConsoleColor_String(t *testing.T) {
	c := NewColor(Red)
	text := c.String("hello world!")
	fmt.Println(text)
}

func TestConsoleColor_Bool(t *testing.T) {
	c1 := ConsoleColor{
		C:           Yellow,
		CancelColor: true,
	}
	p := c1.Bool(true)

	fmt.Printf("%v\n", p)

	c2 := NewColor(White)
	c2.CancelColor = false
	p2 := c2.Bool(true)
	if p != reset+"true" {
		t.Error("can output true")
		return
	}
	fmt.Printf("%v\n", p2)
}
