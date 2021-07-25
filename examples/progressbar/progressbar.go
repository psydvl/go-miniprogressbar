package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/psydvl/goltools/progressbar"
)

func channel(total int) {
	pbinterface, pbwait := progressbar.Init("channel", 0, 0)
	defer pbwait()
	pbch := pbinterface.(chan<- int)
	defer close(pbch)

	pbch <- total
	for i := 0; i < total; i++ {
		pbch <- i
		time.Sleep(time.Millisecond * 400)
	}
	//pbch <- total //alternative variant for ending
}

func step(total int) {
	pbinterface, pbwait := progressbar.Init("step", 0, total)
	defer pbwait()
	pbstep := pbinterface.(func())

	for i := 0; i < total; i++ {
		pbstep()
		time.Sleep(time.Millisecond * 400)
	}
}

func fast(total int) {
	pbinterface, pbwait := progressbar.Init("channel", 0, total)
	defer pbwait()
	pbch := pbinterface.(chan<- int)

	for i := 0; i < total; i++ {
		pbch <- i
		time.Sleep(time.Microsecond * 10) // Not too fast
	}
	close(pbch)

}

func random(total int) {
	pbinterface, pbwait := progressbar.Init("channel", 0, 0)
	defer pbwait()
	pbch := pbinterface.(chan<- int)

	pbch <- total
	for i := 0; i < total; i++ {
		pbch <- i
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2_000)))
	}
	close(pbch)

}

func main() {
	total := 10
	fmt.Println("Running channel variant")
	channel(total)
	fmt.Println("Running step variant")
	step(total)
	fmt.Println("Running channel variant fast (for preview)")
	total = 7000
	fast(total)
	fmt.Println("Running channel variant with random (for preview)")
	total = 70
	random(total)
}
