package lrucache

import (
	"errors"
	"fmt"
)

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}


type lruCache struct {
	cap, length int
	values map[string]*node

	root *node
}

type node struct {
	key string
	value interface{}
	next *node
	prev *node
}

func (n node) String() string {
	return fmt.Sprintf("%v", n.value)
}

func New(cap int) Cache {
	return &lruCache{
		cap: cap,
		values: map[string]*node{},
	}
}

func (l *lruCache) Get(key string) (interface{}, error) {
	val, ok := l.values[key]
	if !ok {
		return nil, errors.New("NotFound")
	}

	ret := val.value
	go l.updateRecentUse(val)
	return ret, nil
}

func (l *lruCache) Set(key string, value interface{}) error {
	newNode := &node{
		key: key,
		value: value,
		next: l.root,
	}
	if l.root != nil {
		l.root.prev = newNode
	}
	l.root = newNode

	_, found := l.values[key]
	if !found {
		l.length++
	}
	l.values[key] = newNode

	if l.length > l.cap {
		cur := l.root
		for i := 0; i < l.cap && cur != nil; i++ {
			cur = cur.next
		}

		if cur != nil {
			cur.prev.next = nil
			for cur != nil {
				delete(l.values, cur.key)
				cur = cur.next
			}
		}
	}
	fmt.Println(l.values)
	return nil
}

func (l *lruCache) updateRecentUse(val *node) {
	if val.prev != nil {
		val.prev.next = val.next
	}
	if val.next != nil {
		val.next.prev = val.prev
	}

	if l.root != nil {
		l.root.prev = val
	}
	val.next = l.root
	val.prev = nil

	l.root = &node{
		key: val.key,
		value: val.value,
		next: val.next,
	}

	cur := l.root
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
}


