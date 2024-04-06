package sisyphus

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func printSlice[T Numeric | string](data []T) {
	for i, value := range data {
		fmt.Print(value)
		// 如果不是最后一个元素，则添加空格
		if i != len(data)-1 {
			fmt.Print(" ")
		}
	}
	// 添加换行符
	fmt.Println()
}

// 统计整数 n 的二进制展开中数位1的总数：O(log n)
func CountOnes[T uint | uint8 | uint16 | uint32 | uint64](n T) int {
	ones := 0
	for 0 < n {
		ones += int(1 & n)
		n >>= 1
	}
	return ones
}

// 数组求和 迭代版本
func SumSliceI(data []int) int {
	sum := 0
	for _, value := range data {
		sum += value
	}
	return sum
}

// 数组求和 递归版本
func SumSlice(data []int, n int) int {
	if n < 1 { // 递归基
		return 0
	} else {
		return SumSlice(data, n-1) + data[n-1]
	}
}

// 数组求和算法（二分递归版，入口为sum(A, 0, n - 1)） O(hi - lo + 1)，线性正比于区间的长度
func SumSlice1B(data []int, lo int, hi int) int {
	if lo == hi {
		return data[lo]
	} else {
		mi := int((lo + hi) >> 1)
		return SumSlice1B(data, lo, mi) + SumSlice1B(data, mi+1, hi)
	}
}

// Fibonacci 数 迭代递归版本
func Fibonacci(n int) uint64 {
	var f, g uint64
	f = 0
	g = 1
	for 0 < n {
		n--
		g += f
		f = g - f
	}
	return f
}

// 置乱器
func UnSortSlice[T Numeric | string](data []T) {
	for i := len(data); i > 0; i-- {
		n, _ := rand.Int(rand.Reader, new(big.Int).SetInt64(int64(i)))
		j := int(n.Int64())
		data[i-1], data[j] = data[j], data[i-1]
	}
	printSlice(data)
}

// 甄别 Slice 是否有序
func DisorderedSlice[T Numeric | string](data []T) int {
	n := 0
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			n++
		}
	}
	return n
}

// 有序 Slice 重复元素剔除
func UniquelySlice[T Numeric | string](data []T) []T {
	i := 0
	for j := 1; j < len(data); j++ {
		if data[i] != data[j] {
			i++
			data[i] = data[j]
		}
	}
	data = data[:i+1]
	printSlice(data)
	return data
}

// 二分查找
func BinSearch[T Numeric | string](data []T, e T) int {
	lo, hi := 0, len(data)
	for 1 < hi-lo {
		mi := (lo + hi) >> 1 // 以中点为轴点
		if e < data[mi] {    // 深入前半[lo, mi)继续查找
			hi = mi
		} else { // 深入后半(mi, hi)继续查找
			lo = mi
		}
	}
	if e == data[lo] {
		return lo
	} else {
		return -1
	}
}

// 欧几里得算法 Euclid algorithm
func Gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
