package quick

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// 分类 ------------ 内部比较排序
// 数据结构 --------- 数组
// 最差时间复杂度 ---- 每次选取的基准都是最大（或最小）的元素，导致每次只划分出了一个分区，需要进行n-1次划分才能结束递归，时间复杂度为O(n^2)
// 最优时间复杂度 ---- 每次选取的基准都是中位数，这样每次都均匀的划分出两个分区，只需要logn次划分就能结束递归，时间复杂度为O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ 主要是递归造成的栈空间的使用(用来保存left和right等局部变量)，取决于递归树的深度，一般为O(logn)，最差为O(n)
// 稳定性 ---------- 不稳定
func QuickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := Partition(a, left, right)
	QuickSort(a, left, pivotIndex-1)
	QuickSort(a, pivotIndex+1, right)
}

func Partition(a []int, left, right int) int {
	pivot := a[right]
	tail := left - 1
	for i := left; i < right; i++ {
		if a[i] <= pivot {
			tail++
			Swap(a, tail, i)
		}
	}
	Swap(a, tail+1, right)
	return tail + 1
}

// written by myself below
func QuickSort2(a []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := Partition2(a, left, right)
	QuickSort2(a, left, pivotIndex-1)
	QuickSort2(a, pivotIndex+1, right)
}

func Partition2(a []int, left, right int) int {
	pivot := a[right]
	tail := left - 1
	for i := left; i < right; i++ {
		if a[i] <= pivot {
			tail++
			Swap(a, tail, i)
		}
	}
	Swap(a, tail+1, right)
	return tail + 1
}
