package cow

import (
	"fmt"
	"github.com/pkg/errors"
	"sync/atomic"
)

type ConcurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}

type concurrentArray struct {
	length uint32
	val    atomic.Value
}

func NewConcurrentArray(length uint32) ConcurrentArray {
	array := concurrentArray{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return &array
}

func (array *concurrentArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}

	newArray := make([]int, array.length)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *concurrentArray) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *concurrentArray) Len() uint32 {
	return array.length
}

func (array *concurrentArray) checkIndex(index uint32) error {
	if index >= array.length {
		return fmt.Errorf("Index out of range [0, %d)!", array.length)
	}
	return nil
}

func (array *concurrentArray) checkValue() error {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}
