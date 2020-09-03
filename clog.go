// Package clog (custom-logger) provides some helpers for quick debugging
// using stdout. Mainly functions that formats a given data and label
// to the output and helpers simplifying the usage of colors.
package clog

import (
	"fmt"
	"log"
)

// Printlb (print-label) prints the input preceded by a label in the format:
//
// My label:
//
// [data shown]
//
// Can be handy for debugging.
func Printlb(v interface{}, label string) {
	log.Printf("%s\n%v\n\n", label, v)
}

// Ok prints any value preceded by a label with a green "OK" tag.
func Ok(v interface{}, label string) {
	Printlb(v, fmt.Sprintf("%s %s", Green("OK"), label))
}

// Error prints any value preceded by a label with a red "Error" tag.
func Error(v interface{}, label string) {
	Printlb(v, fmt.Sprintf("%s %s", Red("Error"), label))
}

// Fatallb (fatal-label) prints the input preceded by a label and exits
func Fatallb(v interface{}, label string) {
	log.Fatalf("%s %s\n%v\n\n", Red("Fatal"), label, v)
}
