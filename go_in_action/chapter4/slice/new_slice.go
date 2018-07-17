package main

import "fmt"

func main() {
	// 初始化 a，长度是3，容量是5。前3个数是10，20，30，后俩数是0，0
	a := make([]int, 3, 5)
	a[0] = 10
	a[1] = 20
	a[2] = 30
	fmt.Println(a, len(a), cap(a))

	// b 是一个长度为2，容量为4的切片，输出的数是 20，30
	b := a[1:3]
	fmt.Println(b, len(b), cap(b))

	// 这里 append 一下，因为 b 容量是够的，所以可以把后面的 0 变为了 1
	b = append(b, 1)
	// 输出是 20，30，1
	fmt.Println(b, len(b), cap(b))
	// 输出还是 10，20，30，但是后面两个数变成了 1，0
	fmt.Println(a, len(a), cap(a))

	// 这里 a 再 append 一下，a 的容量也是够的，所以把后面的 1 变成了 2
	a = append(a, 2)
	// 这里输出10，20，30，2，长度是4，容量还是5
	fmt.Println(a, len(a), cap(a))
	// ！！！！！注意，这里 b 的底层数组来自 a，所以 b 本来是20，30，1，现在变成了 20，30，2，长度3，容量4
	fmt.Println(b, len(b), cap(b))

	// c 是20，30，2，长度3，容量也是3
	c := a[1:4:4]
	fmt.Println(c, len(c), cap(c))
	// append 的时候，发现容量不够了，所以会重新分一个区域，复制出一个长度3，容量6的空间，再append，就变成了长度4，容量6
	c = append(c, 100)
	// 输出是 20，30，2，100
	fmt.Println(c, len(c), cap(c))

	a[2] = 300
	// 这里a是 10，20，300，2；b 是20，300，2；c 是20，30，2，100
	// 因为 b 和 a还是共享同一个底层的数组，所以修改a，也会影响b
	// 但是 c 已经是复制出来了，不再和a、b共享同一个底层的数组了，所以不回变化
	fmt.Println(a, b, c)
}
