package bucket

import (
	"fmt"
	"sort"
)

const (
	bucketNumber = 5
)

var C []int

func InsertionSort(a []int, left, right int) {
	for i := left + 1; i <= right; i++ {
		get := a[i]
		j := i - 1
		for j >= left && a[j] > get {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = get
	}
}

func MapToBucket(x int) int {
	return x / 10
}

func CountingSort(a []int, n int) {
	C = make([]int, bucketNumber, bucketNumber)
	for i := 0; i < n; i++ {
		C[MapToBucket(a[i])] += 1
	}
	for i := 1; i < bucketNumber; i++ {
		C[i] += C[i-1]
	}
	fmt.Println(C)

	b := make([]int, n, n)
	for i := n - 1; i >= 0; i-- {
		tmp := MapToBucket(a[i])
		C[tmp] -= 1
		b[C[tmp]] = a[i]
	}
	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
	sort.Sort()
}

// 分类 ------------- 内部非比较排序
// 数据结构 --------- 数组
// 最差时间复杂度 ---- O(nlogn)或O(n^2)，只有一个桶，取决于桶内排序方式
// 最优时间复杂度 ---- O(n)，每个元素占一个桶
// 平均时间复杂度 ---- O(n)，保证各个桶内元素个数均匀即可
// 所需辅助空间 ------ O(n + bn)
// 稳定性 ----------- 稳定
func BucketSort(a []int, n int) {
	CountingSort(a, n)
	// CountingSort 之后，各个桶都放好了该放的数字，但是每个桶内的顺序还没有排。
	// 下面就是给每个桶进行插入排序。
	// 关键在于找到每个桶的左右边界：这里利用了辅助数组C。
	// CountingSort之后，C从 [2 2 5 6 8] 变成了 [0 2 2 5 6]
	// 所以第 i 个桶的左边界是 C[i]，右边界是 C[i+1]-1，注意最后一个桶的右边界只能是 n-1
	for i := 0; i < bucketNumber; i++ {
		left := C[i]
		var right int
		if i == bucketNumber-1 {
			right = n - 1
		} else {
			right = C[i+1] - 1
		}
		if left < right {
			InsertionSort(a, left, right)
		}
	}
}
