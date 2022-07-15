package go_console_color

import (
	"fmt"
	"testing"
)

func TestConsoleColor_String(t *testing.T) {
	c := ConsoleColor{
		C:           Red,
		CancelColor: true,
	}
	p := c.String("spring")
	if p != reset+"spring" {
		t.Error("can output spring")
		return
	}
	fmt.Printf("%v\n", p)
}

func TestConsoleColor_Bool(t *testing.T) {
	c1 := ConsoleColor{
		C:           Yellow,
		CancelColor: true,
	}
	p := c1.Bool(true)

	fmt.Printf("%v\n", p)

	c2 := ConsoleColor{
		C:           Yellow,
		CancelColor: false,
	}
	p2 := c2.Bool(true)
	if p != reset+"true" {
		t.Error("can output true")
		return
	}
	fmt.Printf("%v\n", p2)
}
