package counting

const k = 13

var C []int

// 分类 ------------ 内部非比较排序
// 数据结构 --------- 数组
// 最差时间复杂度 ---- O(n + k)
// 最优时间复杂度 ---- O(n + k)
// 平均时间复杂度 ---- O(n + k)
// 所需辅助空间 ------ O(n + k)
// 稳定性 ----------- 稳定
func CountingSort(a []int) {
	C = make([]int, k, k)
	length := len(a)
	for i := 0; i < length; i++ {
		C[a[i]] = C[a[i]] + 1
	}
	for i := 1; i < k; i++ {
		C[i] = C[i] + C[i-1]
	}
	b := make([]int, length, length)
	for i := length - 1; i >= 0; i-- {
		C[a[i]] -= 1
		b[C[a[i]]] = a[i]
	}
	for i := 0; i < length; i++ {
		a[i] = b[i]
	}
}

// written by myself below
func CountingSort2(a []int) {
	C = make([]int, k, k)
	length := len(a)
	for i := 0; i < length; i++ {
		C[a[i]] += 1
	}
	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}
	b := make([]int, length, length)
	for i := length - 1; i >= 0; i-- {
		C[a[i]] -= 1
		b[C[a[i]]] = a[i]
	}
	for i := 0; i < length; i++ {
		a[i] = b[i]
	}
}
