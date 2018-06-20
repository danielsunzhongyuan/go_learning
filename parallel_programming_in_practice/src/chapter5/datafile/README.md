# v1版本
v1版本就是利用了互斥锁和读写锁实现了文件操作的并发安全。
但是也存在一个问题，就是在 Read方法里，需要添加一个for循环，防止读出空的数据。
这里的 df.fmutex.RUnlock() 很多，不是很好。

# v2版本
v2版本引入了条件变量 - Cond 来解决这个问题。
需要对Read和Write都做一些改造。

# v3版本
v3版本引入了原子操作，来减少并发安全带来的性能损耗。

原子操作即不能被中断的操作。Go的标准库 sync/atomic 提供了众多的函数代表。
这些函数可以对集中简单类型的值执行原子操作。
类型有6种：int32, int64, uint32, uint64, uintptr, unsafe.Pointer
操作有5种：增/减，比较并交换，载入，存储，交换。

### 增/减
加3：
```
atomic.AddInt32(&i32, 3)
atomic.AddInt64(&i64, 3)
atomic.AddUint32(&ui32, 3)
atomic.AddUint64(&ui64, 3)
atomic.AddUintptr(&uintptr, 3)
```
减3：
```
atomic.AddInt32(&i32, -3)
atomic.AddInt64(&i64, -3)
atomic.AddUint32(&i32, ^uint32(-NN-3))
atomic.AddUint64(&i64, ^uint64(-NN-3))
atomic.AddUintptr(&uintptr, ^uintptr(-NN-3))
```

这是因为后面的这个参数要和前面的参数类型一致，所以利用了反码的特性。

unsafe.Pointer无法加减。

### 比较并交换（CAS，Compare and Swap）
```Go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
```
3个参数，被操作值的指针，被操作值的旧值和新值。如果被操作值等于旧值，则替换为新值，swapped为true。
CAS的优势是：在不创建互斥量和不形成临界区的情况下，完成并发安全的值替换操作。可以优化性能。
劣势是：如果被操作值频繁变更，CAS操作不那么容易成功。所以需要for循环来进行多次尝试。
```go
var value int32
func addValue(delta int32) {
    for {
        v := value
        if atomic.ComapareAndSwapInt32(&value, v, (v+delta)) {
            break
        }
    }
}
```

CompareAndSwapInt32, CompareAndSwapInt64, CompareAndSwapPointer
CompareAndSwapUint32, CompareAndSwapUint64, CompareAndSwapUintptr.

**应该优先选择使用CAS**

### 载入 Load
为了原子的读取某个值，引入了载入。
```go
var value int32
func addValue(delta int32) {
    for {
        v := atomic.LoadInt32(&value) // 即原子地读取value并赋值给 v
        if atomic.ComapareAndSwapInt32(&value, v, (v+delta)) {
            break
        }
    }
}
```

atomic.LoadInt32, atomic.LoadInt64, atomic.LoadPointer,
atomic.LoadUint32, atomic.LoadUint64, atomic.LoadUintptr.

### 存储 Store
与读取操作对应的是写入操作。

atomic.StoreInt32, atomic.StoreInt64, atomic.StorePointer,
atomic.StoreUint32, atomic.StoreUint64, atomic.StoreUintptr.

### 交换 Swap
Swap不关心旧值，而是直接交换。比CAS约束少，比Load操作功能强。
atomic.SwapInt32 的第一个参数是*int32，第二个表示新值，返回的是旧值。

### 原子值 Value
atomic.Value是一个结构体类型，用于存储需要原子读写的值。它可以接受的被操作值的类型不限。
```
var atomicVal atomic.Value
```
只有两个方法，Load和Store。
Load无参数，返回一个 interface{}
Store参数是 interface{}，没有返回结果。

Store的限制：
1. 传入的参数不能为 nil
2. 传入的参数的类型必须和以前传进来的类型一致。也就是不能变。
