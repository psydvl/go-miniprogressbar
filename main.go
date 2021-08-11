package goltools

import (
	"log"

	"github.com/psydvl/goltools/progressbar"
	"github.com/psydvl/goltools/terminal"
)


func init() {
	log.SetPrefix("goltools: ")
	log.SetFlags(0)
	log.Println("github.com/psydvl/goltools is initialized")
}

func ProgessBar(total int) chan<- int {
	var ch interface{}
	if total <= 0 {
		total = 0
	}
	ch, _ = progressbar.Init("channel", 0, total)
	return ch.(chan<- int)
}

func TerminalWidth() int {
	return terminal.Width()
}