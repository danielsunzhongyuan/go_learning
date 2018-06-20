package radix

const (
	dn = 3
	k  = 10
)

var C []int

// 分类 ------------- 内部非比较排序
// 数据结构 ---------- 数组
// 最差时间复杂度 ---- O(n * dn) dn表示待排序元素的数字位数
// 最优时间复杂度 ---- O(n * dn)
// 平均时间复杂度 ---- O(n * dn)
// 所需辅助空间 ------ O(n * dn)
// 稳定性 ----------- 稳定
func LsdRadixSort(a []int, n int) {
	for d := 1; d <= dn; d++ {
		CountingSort(a, n, d)
	}
}

func CountingSort(a []int, n, d int) {
	C := make([]int, k, k)
	for i := 0; i < n; i++ {
		C[GetDigit(a[i], d)] += 1
	}
	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}
	b := make([]int, n, n)
	for i := n - 1; i >= 0; i-- {
		digit := GetDigit(a[i], d)
		C[digit] -= 1
		b[C[digit]] = a[i]
	}
	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
}

func GetDigit(x, d int) int {
	radix := []int{1, 1, 10, 100}
	return (x / radix[d]) % 10
}

// written by myself below
func LsdRadixSort2(a []int, n int) {
	for d := 1; d <= dn; d++ {
		CountingSort2(a, n, d)
	}
}

func CountingSort2(a []int, n, d int) {
	C := make([]int, k, k)
	for i := 0; i < n; i++ {
		C[GetDigit(a[i], d)] += 1
	}
	for i := 1; i < k; i++ {
		C[i] += C[i-1]
	}

	b := make([]int, n, n)
	for i := n - 1; i >= 0; i-- {
		digit := GetDigit(a[i], d)
		C[digit] -= 1
		b[C[digit]] = a[i]
	}
	for i := 0; i < n; i++ {
		a[i] = b[i]
	}
}
