package main

import (
	"fmt"
	"sync"
	"time"
)

type LockeAbleNotEmbeded struct {
	lock           sync.Mutex
	readableInt    int
	readableString string
}

func main() {
	l := &LockeAbleNotEmbeded{}
	go func() {
		for {
			time.Sleep(time.Second * 15)
			l.read("one")
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 10)
			l.write("two", "hello", 5)
		}
	}()

	l.lock.Lock()
	l.lock.Unlock()
	l.read("one")

	select {}
}

func (l *LockeAbleNotEmbeded) read(st string) {
	l.lock.Lock()
	fmt.Println("just Locked by", st, "for reading")
	fmt.Println(l.readableInt, l.readableString)
	l.lock.Unlock()
	fmt.Println("just UnLocked", st, "from reading")
}

func (l *LockeAbleNotEmbeded) write(st string, val string, valInt int) {
	l.lock.Lock()
	fmt.Println("just Locked by", st, "for writing")
	l.readableString = val
	l.readableInt = valInt
	l.lock.Unlock()
	fmt.Println("just UnLocked", st, "from writing")
}
