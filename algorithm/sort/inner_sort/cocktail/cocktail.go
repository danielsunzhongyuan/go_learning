package cocktail

// 分类 -------------- 内部比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n^2)
// 最优时间复杂度 ---- 如果序列在一开始已经大部分排序过的话,会接近O(n)
// 平均时间复杂度 ---- O(n^2)
// 所需辅助空间 ------ O(1)
// 稳定性 ------------ 稳定
func CockTailSort(a []int) {
	length := len(a)
	left, right := 0, length-1
	for left < right {
		for i := left; i < right; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		right--
		for i := right; i > left; i-- {
			if a[i] < a[i-1] {
				a[i], a[i-1] = a[i-1], a[i]
			}
		}
		left++
	}
}

func CockTailSort2(a []int) {
	length := len(a)
	left, right := 0, length-1
	for left < right {
		for i := left; i < right; i++ {
			SwapIfNeeded(a, i, i+1)
		}
		right--
		for j := right; j > left; j-- {
			SwapIfNeeded(a, j-1, j)
		}
		left++
	}
}

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
func SwapIfNeeded(a []int, i, j int) {
	if a[i] > a[j] {
		Swap(a, i, j)
	}
}
