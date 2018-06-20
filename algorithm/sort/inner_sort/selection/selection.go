package selection

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ----- O(n^2)
// 最优时间复杂度 ----- O(n^2)
// 平均时间复杂度 ----- O(n^2)
// 所需辅助空间 ------- O(1)
// 稳定性 ------------ 不稳定
func SelectionSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			a[i], a[min] = a[min], a[i]
		}
	}
}

func SelectionSort2(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		if min != i {
			Swap(a, min, i)
		}
	}
}

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
