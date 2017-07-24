package main

import (
	"fmt"
	"time"

	"github.com/sadlil/experiments/go/broadcast/simplebroadcaster"
	"github.com/sadlil/experiments/go/broadcast/complexbroadcaster"
)

const (
	colorGreen = iota + 32
)

const escape = "\x1b"

func main() {
	t := time.Now()
	simpleBroadcast()
	fmt.Printf("%s[1;32mSimple Broadcaster took %s\n", escape, time.Since(t).String())
	fmt.Printf("%s[0m", escape)

	t = time.Now()
	complexBroadcaster()
	fmt.Printf("%s[1;32mComplex Broadcaster took %s\n", escape, time.Since(t).String())
	fmt.Printf("%s[0m", escape)
}

func simpleBroadcast() {
	br := simplebroadcaster.NewBroadcaster()

	for i := 1; i <= 10; i++ {
		list := br.Listen()
		go func(i int, ch chan interface{}) {
			select {
			case val := <-ch:
				fmt.Println("Routine ", i, "Recived value", val)
			}
		}(i, list)
	}

	br.Broadcast("[simple] Hello World")
	time.Sleep(3 * 1e9)
}

func complexBroadcaster() {
	br := complexbroadcaster.NewBroadcaster()
	for i := 1; i <= 10; i++ {
		list := br.Listen()
		go func(i int, r complexbroadcaster.Receiver) {
			for v := r.Read(); v != nil; v = r.Read() {
				fmt.Println("Routine ", i, "Recived value", v)
			}
		}(i, list)
	}

	br.Broadcast("[complex] Hello World")
	time.Sleep(3 * 1e9)
}