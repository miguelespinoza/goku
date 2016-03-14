package main

import (
	"fmt"
	"time"

	"github.com/miguelespinoza/goku/goku"
)

func main() {
	puzzle := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"

	start := time.Now()
	fmt.Println("goku solving...")
	output, err := goku.Solve(puzzle)
	if err != nil {
		fmt.Println(err)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	goku.PrettyDisplay(output)
}
