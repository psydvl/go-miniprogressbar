package mpb

import (
	"fmt"
	"strings"
	"time"
)

// TODO Maybe can reuse or should delete?
type ProgressBar struct {
	ch <-chan int
	length int
	total int
}

// output example
// [■■■■■■■■■■■■■■■■■                                                     ] 1248/5000  6s spent
func pb(progress ProgressBar) {
	var ch = progress.ch
	var length = progress.length
	if progress.total == 0 {
		progress.total = <-ch
	}
	var total = progress.total

	var start time.Time = time.Now()
	var bar int
	for current := range ch {
		bar = length * current / total
		fmt.Printf("\r\033[1;34m[%s%s] %d/%d  %v spent\033[0m   ",
			strings.Repeat("■", bar),
			strings.Repeat(" ", length-bar),
			current,
			total,
			time.Since(start).Round(time.Second))
	}
	fmt.Printf("\r\033[1;34m[%s] %d/%d  %v spent\033[0m   \n",
		strings.Repeat("■", length),
		total, total,
		time.Since(start).Round(time.Second))
}

func steps(ch_step <-chan bool, ch_out chan<- int) {
	var current int = 0
	for range ch_step {
		current += 1
		ch_out <- current
	}
	close(ch_out)
}

/*
	Required bar length and total progress steps as input
*/
func Main(length int, total int) chan<- int {
	var ch chan int = make(chan int)
	var progress ProgressBar
	if length == 0 {
		length = 70
	}
	progress = ProgressBar{ch, total, length}
	go pb(progress)
	return ch
}

/*
	Required bar length as input;
	First value in ouputed channel should be total progress steps
*/
func Simple(length int) chan<- int {
	var ch chan int = make(chan int)
	if length == 0 {
		length = 70
	}
	var progress ProgressBar = ProgressBar{ch, 0, length}
	go pb(progress)
	return ch
}

/*
	Required bar length and total progress steps as input;
	Any value in ouputed channel will increase progress
*/
func Steps(length int, total int) chan<- bool {
	var ch_steps chan bool = make(chan bool)
	var ch chan int = make(chan int)
	var progress ProgressBar
	if length == 0 {
		length = 70
	}
	if total <= 0 {
		panic("Zero or negative total value provided")
	}
	progress = ProgressBar{ch, total, length}
	go pb(progress)
	go steps(ch_steps, ch)
	return ch_steps
}