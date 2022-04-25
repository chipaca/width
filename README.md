# width
The width package helps you manage string widths

When writing code that wants to print things to a terminal while
formatting those things using escape sequences, often you want to keep
track of the width of the thing you've printed, so you can do things
like word breaking, line wrapping, or alignment.

The `width` package aims to be the basic block for solving that family
of problems.

At its core there's just an interface,

```go
type StringWidther interface {
	String() string
	Width() int
}
```

and two implementations of that thing: a bare `String` type that just uses
[go-runewidth](https://github.com/mattn/go-runewidth), and an explicit `StringAndWidth`
that lets you set both things by hand. This is enough for most things.

---

At some point this library might go beyond that looking at actual widths for
some currently poorly-supported corner cases (i.e. cases where runewidth and
your terminal disagree). So far just ingoring that issue has solved most
problems (e.g. VTE in Ubuntu 16.04 mangles emoji widths, but nobody cares about
it any more).
