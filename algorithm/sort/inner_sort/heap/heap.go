package heap

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Heapify(a []int, i, size int) {
	leftChild := 2*i + 1
	rightChild := 2*i + 2
	max := i
	if leftChild < size && a[leftChild] > a[max] {
		max = leftChild
	}
	if rightChild < size && a[rightChild] > a[max] {
		max = rightChild
	}
	if max != i {
		Swap(a, i, max)
		Heapify(a, max, size)
	}
}

func BuildHeap(a []int, n int) int {
	heapSize := n
	for i := heapSize/2 - 1; i >= 0; i-- {
		Heapify(a, i, heapSize)
	}
	return heapSize
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(nlogn)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 不稳定
func HeapSort(a []int, n int) {
	heapSize := BuildHeap(a, n)
	for heapSize > 1 {
		heapSize--
		Swap(a, 0, heapSize)
		Heapify(a, 0, heapSize)
	}
}

// written by myself
func HeapSort2(a []int, n int) {
	heapSize := BuildHeap2(a, n)
	for heapSize > 1 {
		heapSize--
		Swap(a, 0, heapSize)
		Heapify(a, 0, heapSize)
	}
}

func BuildHeap2(a []int, n int) int {
	heapSize := n
	for i := heapSize/2 - 1; i >= 0; i-- {
		Heapify2(a, i, heapSize)
	}
	return heapSize
}

func Heapify2(a []int, i, size int) {
	left := 2*i + 1
	right := 2*i + 2
	max := i
	if left < size && a[max] < a[left] {
		max = left
	}
	if right < size && a[max] < a[right] {
		max = right
	}
	if max != i {
		Swap(a, i, max)
		Heapify2(a, max, size)
	}
}
