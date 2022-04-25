// Package width defines an interface, `StringWidther`, and a couple
// of trivial implementations of that interface, to allow you to build
// things where the length of a string and its width on screen are
// independent of one another. This is useful for embedding escape
// sequences, for example, as well as for accounting for glyphs of
// varying on-screen width.
package width

import (
	"github.com/mattn/go-runewidth"
)

// The StringWidther interface is the basic interface this package
// builds on.
type StringWidther interface {
	String() string
	Width() int
}

// A String is just a string that uses runewidth to get its width.
type String string

func (s String) String() string {
	return string(s)
}

// Width of a String is just its runewidth.StringWidth (ie the sum of
// the widths of its runes).
func (s String) Width() int {
	return runewidth.StringWidth(string(s))
}

// StringAndWidth keeps track of its string and its width independent
// and explicitly.
type StringAndWidth struct {
	S string
	W int
}

func (sw StringAndWidth) String() string {
	return sw.S
}

// Width of a StringAndWidth is whatever it was set to.
func (sw StringAndWidth) Width() int {
	return sw.W
}
