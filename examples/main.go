package main

import (
	"time"

	mpb "github.com/psydvl/go-miniprogressbar"
)

func main() {
	ch := mpb.Main(0, 0)
	ch <- 10
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
	time.Sleep(time.Second)
}
