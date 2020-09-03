package clog

import (
	"runtime"
)

type color struct {
	reset string
	red   string
	green string
	blue  string
}

var c = color{}

func init() {
	// Not compatible with Windows
	if runtime.GOOS != "windows" {
		c.reset = "\033[0m"
		c.red = "\033[31m"
		c.green = "\033[32m"
		c.blue = "\033[34m"
	}
}

// Red returns the input string in red on compatible systems
func Red(s string) string {
	return c.red + s + c.reset
}

// Green returns the input string in green on compatible systems
func Green(s string) string {
	return c.green + s + c.reset
}

// Blue returns the input string in blue on compatible systems
func Blue(s string) string {
	return c.blue + s + c.reset
}
