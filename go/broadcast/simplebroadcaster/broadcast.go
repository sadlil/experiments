package simplebroadcaster

type Broadcaster struct {
	listeners []chan interface{}
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{
		listeners: make([]chan interface{}, 0),
	}
}

func (b *Broadcaster) Broadcast(v interface{}) {
	for _, listener := range b.listeners {
		listener <- v
	}
}

func (b *Broadcaster) Listen() chan interface{} {
	newCh := make(chan interface{})
	b.listeners = append(b.listeners, newCh)
	return newCh
}

