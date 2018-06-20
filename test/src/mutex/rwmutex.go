package main

import (
	"sync"
	"time"
)

var rwMutex *sync.RWMutex

/*
读写锁是针对读写的互斥锁

基本遵循两大原则：

1、可以随便读，多个goroutine同时读

2、写的时候，啥也不能干。不能读也不能写

RWMutex提供了四个方法：

func (*RWMutex) Lock // 写锁定
func (*RWMutex) Unlock // 写解锁
func (*RWMutex) RLock // 读锁定
func (*RWMutex) RUnlock // 读解锁
*/

func main() {
	rwMutex = new(sync.RWMutex)

	// 写的时候啥也不能干
	go write(1)
	go write(2)
	go read(3)
	go read(4)
	go read(5)
	go write(6)

	time.Sleep(6 * time.Second)
}

func read(i int) {
	println(i, "read start")

	rwMutex.RLock()
	defer rwMutex.RUnlock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read over")

}

func write(i int) {
	println(i, "write start")

	rwMutex.Lock()
	defer rwMutex.Unlock()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write over")

}
