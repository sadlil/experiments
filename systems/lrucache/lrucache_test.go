package lrucache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	c := New(2)

	_, err := c.Get("notfound")
	assert.NotNil(t, err)

	_ = c.Set("one", 1)
	_, err = c.Get("one")
	assert.Nil(t, err)

	_ = c.Set("two", 2)
	_ = c.Set("three", 3)
	_ = c.Set("four", 4)

	_, err = c.Get("one")
	assert.NotNil(t, err)
}
