package cmap

import (
	"sync"
	"sync/atomic"
)

type Bucket interface {
	Put(p Pair, lock sync.Locker) (bool, error)
	Get(key string) Pair
	GetFirstPair() Pair
	Delete(key string, lock sync.Locker) bool
	Clear()
	Size() uint64
	String() string
}

type bucket struct {
	firstValue atomic.Value
	size       uint64
}

var placeholder = &pair{}

func newBucket() Bucket {
	b := &bucket{}
	b.firstValue.Store(placeholder)
	return b
}

func (b *bucket) Put(p Pair, lock sync.Locker) (bool, error) {

}
