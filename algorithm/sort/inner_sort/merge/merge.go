package merge

import "fmt"

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(nlogn)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(nlogn)
// 所需辅助空间 ------ O(n)
// 稳定性 ------------ 稳定
func Merge(a []int, left, mid, right int) {
	length := right - left + 1
	tmp := make([]int, length, length)
	index := 0
	i := left
	j := mid + 1
	for i <= mid && j <= right {
		if a[i] <= a[j] {
			tmp[index] = a[i]
			i++
			index++
		} else {
			tmp[index] = a[j]
			j++
			index++
		}

	}
	for i <= mid {
		tmp[index] = a[i]
		i++
		index++
	}
	for j <= right {
		tmp[index] = a[j]
		index++
		j++
	}

	for k := 0; k < length; k++ {
		a[left] = tmp[k]
		left++
	}
}

func MergeSortRecursion(a []int, left, right int) {
	fmt.Println("left:", left, ",right:", right)
	if left >= right {
		return
	}
	mid := (left + right) / 2
	MergeSortRecursion(a, left, mid)
	MergeSortRecursion(a, mid+1, right)
	Merge(a, left, mid, right)
}

func MergeSortIteration(a []int, length int) {
	a = append(a)
	var left, mid, right int
	for i := 1; i < length; i *= 2 {
		left = 0
		for left+i < length {
			mid = left + i - 1
			if mid+i < length {
				right = mid + i
			} else {
				right = length - 1
			}
			Merge(a, left, mid, right)
			left = right + 1
		}
	}
}
