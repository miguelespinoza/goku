package bridge

import "github.com/miguelespinoza/goku/goku"

// The mobile package provides a very narrow interface to ivy,
// suitable for wrapping in a UI for mobile applications.
// It is designed to work well with the gomobile tool by exposing
// only primitive types. It's also handy for testing.

// Solve : proxy function so that Java/Objective-C binding can communicate with Golang natively
func Solve(unsolved string) (result string, errors error) {
	return goku.SolveDirect(unsolved)
}
