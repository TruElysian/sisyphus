package sisyphus

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func PrintSlice[T Comparable](items []T, str string) {
	for i, value := range items {
		fmt.Print(value)
		// 如果不是最后一个元素，则添加空格
		if i != len(items)-1 {
			fmt.Print(str)
		}
	}
	// 添加换行符
	fmt.Println()
}

// MARK: Slice 相关函数

// 根据 index 删除 Slice 中的对应元素
func RemoveSlice[T Comparable](items []T, index int) ([]T, error) {
	if index < 0 || index >= len(items) {
		return items, fmt.Errorf("index out of range")
	}
	return append(items[:index], items[index+1:]...), nil
}

// 置乱器
func UnSortSlice[T Numeric | string](items []T) {
	for i := len(items); i > 0; i-- {
		n, _ := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(i)))
		j := int(n.Int64())
		items[i-1], items[j] = items[j], items[i-1]
	}
	PrintSlice(items, " ")
}

// 甄别 Slice 是否有序 返回 Slice 中逆序相邻元素对的总数
func DisorderedSlice[T Numeric | string](items []T) int {
	n := 0
	for i := 1; i < len(items); i++ {
		if items[i-1] > items[i] {
			n++
		}
	}
	return n
}

// 无序 Slice 重复元素剔除
func DeduplicateSlice[T Comparable](items []T) []T {
	seen := make(map[T]bool) // 借助 Hash Map
	var result []T
	for _, value := range items {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

// 有序 Slice 重复元素剔除
func UniquelySlice[T Comparable](items []T) []T {
	i := 0
	for j := 1; j < len(items); j++ {
		if items[i] != items[j] {
			i++
			items[i] = items[j] //发现不同元素时，向前移至紧邻于前者右侧
		}
	}
	items = items[:i+1] // 截取多余的重复元素
	return items
}

// MARK: Fibonacci

// Fibonacci 数 迭代版本
func Fibonacci(n int) uint64 {
	f, g := uint64(0), uint64(1)
	for 0 < n {
		n--
		g += f
		f = g - f
	}
	return f
}

// 返回 Fibonacci 的 Slice
func FibonacciSlice(n int) []uint64 {
	var data []uint64
	f, g := uint64(0), uint64(1)
	for 0 < n {
		n--
		g += f
		f = g - f
		data = append(data, f)
	}
	return data
}

// MARK: 查找算法

// 二分查找 O(logn)
func BinSearch[T Comparable](items []T, e T) int {
	lo, hi := 0, len(items)
	for 1 < hi-lo {
		mi := (lo + hi) >> 1 // 以中点为轴点
		if e < items[mi] {   // 深入前半 [lo, mi) 继续查找
			hi = mi
		} else { // 深入后半 (mi, hi) 继续查找
			lo = mi
		}
	} // 出口时 hi=lo+1，查找区间仅包含一个元素 data[lo]
	if e == items[lo] {
		return lo
	} else {
		return -1
	}
}

func BinSearch1[T Comparable](items []T, e T) int {
	lo, hi := 0, len(items)
	for lo < hi {
		mi := (lo + hi) >> 1 // 以中点为轴点
		if e < items[mi] {   // 深入前半 [lo, mi) 继续查找
			hi = mi
		} else { // 深入后半 (mi, hi) 继续查找
			lo = mi + 1
		}
	} // 成功查找不能提前终止
	lo--
	return lo
}

// MARK: 排序算法

// 归并排序
func MergeSort[T Comparable](items []T) []T {
	if len(items) < 2 { // 单元素区间自然有序
		return items
	}
	mi := len(items) >> 1
	left := MergeSort(items[:mi])  // [lo,mi)
	right := MergeSort(items[mi:]) // [mi,hi)
	return merge(left, right)
}

func merge[T Comparable](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素追加到结果数组
	if i < len(left) {
		result = append(result, left[i:]...)
	}
	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

// MARK: Others

// 欧几里得算法 Euclid algorithm
func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}

// 快速幂 迭代版本
func BinPow(base, exponent uint64) uint64 {
	result := uint64(1) // 初始化结果为 1
	for exponent > 0 {
		if exponent&1 != 0 { // 如果指数是奇数
			result *= base // 更新结果
		}
		base *= base   // 底数平方
		exponent >>= 1 // 右移指数
	}
	return result
}

// 统计整数 n 的二进制展开中数位 1 的总数：O(log n)
func CountOnes[T uint | uint8 | uint16 | uint32 | uint64](n T) int {
	ones := 0
	for 0 < n {
		ones += int(1 & n)
		n >>= 1
	}
	return ones
}
