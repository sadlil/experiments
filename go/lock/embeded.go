package main

import (
	"fmt"
	"sync"
	"time"
)

type LockeAble struct {
	sync.Mutex
	readableInt    int
	readableString string
}

func main() {
	l := &LockeAble{}
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

	l.Lock()
	l.Unlock()
	l.read("one")

	select {}
}

func (l *LockeAble) read(st string) {
	l.Lock()
	fmt.Println("just Locked by", st, "for reading")
	fmt.Println(l.readableInt, l.readableString)
	l.Unlock()
	fmt.Println("just UnLocked", st, "from reading")
}

func (l *LockeAble) write(st string, val string, valInt int) {
	l.Lock()
	fmt.Println("just Locked by", st, "for writing")
	l.readableString = val
	l.readableInt = valInt
	l.Unlock()
	fmt.Println("just UnLocked", st, "from writing")
}
