package main

// Node represents a cell in the grid.
// Fields are exported so it can be used from other packages.
type Node struct {
	X   int
	Y   int
	Val rune
}
