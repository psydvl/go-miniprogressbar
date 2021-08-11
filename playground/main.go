package main

import (
	"fmt"

	"github.com/psydvl/goltools/terminal"
)

func main() {
	result := terminal.Width()
	fmt.Println(result)
}