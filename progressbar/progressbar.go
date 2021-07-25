package progressbar

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/psydvl/goltools"
)

type progressBar struct {
	wg *sync.WaitGroup
	ch <-chan int
	length int
	total int
}

type progressStep struct {
	ch chan<- int
	current int
}

// output example
//
// [■■■■■■■■■■■■■■■■■                                                     ] 1248/5000  6s spent
func (progress progressBar)iterate() {
	var ch = progress.ch
	var length = progress.length
	if progress.total == 0 {
		progress.total = <-ch
		if progress.total <= 0 {
			panic("Wrong progress total value from channel")
		}
	}
	var total = progress.total
	defer progress.wg.Done()
	fmt.Printf("\r\033[1;34m[%s] 0/%d  0s spent\033[0m   ",
		strings.Repeat(" ", length),
		total,
	)

	var start time.Time = time.Now()
	var bar int
	var duration time.Duration
	for current := range ch {
		duration = time.Since(start).Round(time.Second)
		if current == progress.total {
			break
		}
		bar = length * current / total
		fmt.Printf("\r\033[1;34m[%s%s] %d/%d  %v spent\033[0m   ",
			strings.Repeat("■", bar),
			strings.Repeat(" ", length-bar),
			current,
			total,
			duration,
		)
	}
	fmt.Printf("\r\033[1;34m[%s] %d/%d  %v spent\033[0m   \n",
		strings.Repeat("■", length),
		total, total,
		duration,
	)
}

func (progress *progressStep)step() {
	progress.current = progress.current + 1
	progress.ch <- progress.current
}

/*
 Initialize progress bar;
 possible methods: "channel", "step"
 
 if length is 0 it will be changed via "tput cols" or set as default '70'

 "channel" returns input int channel as empty interface and sync.WaitGroup.Wait

 "step" returns function as empty interface and sync.WaitGroup.Wait
 */
func Init(method string, length, total int) (interface{}, func()) {
	var result interface{}

	var ch chan int = make(chan int)
	var wg sync.WaitGroup
	var progress progressBar

	if length == 0 {
		length = (goltools.TerminalWidth() - 27) / 10 * 10
	}
	if length == 0 {
		length = 70
	}

	progress = progressBar{&wg, ch, length, total}

	switch method {
	case "channel":
		if progress.total < 0 {
			panic(fmt.Sprintf("total variable is negative, got: %d", progress.total))
		}
		var iter chan<- int = ch
		wg.Add(1)
		go progress.iterate()
		result = iter
	case "step":
		if progress.total <= 0 {
			panic(fmt.Sprintf("total variable is negative or zero, got: %d", progress.total))
		}
		step := progressStep{ch, 0}
		wg.Add(1)
		go progress.iterate()
		result = step.step
	default:
		panic("progressbar: wrong method value; \"step\" or \"channel\" expected")
	}
	
	return result, wg.Wait
}

//Created input channel for progressbar;
//if total is 0 will be rewrited with first channel value
func Simple(length, total int) chan<- int {
	ch, _ := Init("channel", length, total)
	return ch.(chan<- int)
}

//Created step function for progressbar
func Step(length, total int) func() {
	iter, _ := Init("step", length, total)
	return iter.(func())
}
