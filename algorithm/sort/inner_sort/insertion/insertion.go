package insertion

import "fmt"

// 分类 ------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- 最坏情况为输入序列是降序排列的,此时时间复杂度O(n^2)
// 最优时间复杂度 ---- 最好情况为输入序列是升序排列的,此时时间复杂度O(n)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定

func InsertionSort(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		tmp := a[i]
		j := i - 1
		// 这里的一个关键是 j >= 0，而不是 j>0
		for j >= 0 && a[j] > tmp {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = tmp
		fmt.Println(a)
	}
}

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- O(nlogn)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定

// 二分插入排序只是减少了比较的次数，对于数据交换的次数并没有减少
// 时间复杂度也没有提升
func InsertionSortDichotomy(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		tmp := a[i]
		left, right := 0, i-1
		for left <= right {
			mid := (left + right) / 2
			if a[mid] > tmp {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := i - 1; j >= left; j-- {
			a[j+1] = a[j]
		}
		a[left] = tmp
		fmt.Println("left:", left)
	}
}
