package main

import (
	"fmt"

	"github.com/psydvl/goltools"
)

func main() {
	result := goltools.TerminalWidth()
	fmt.Println(result)
}