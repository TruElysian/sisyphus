package sisyphus_test

import (
	"fmt"
	"sisyphus"
	"testing"
)

func TestRemoveSlice(t *testing.T) {
	data := []int{1, 4, 1, 2, 2, 5, 3, 4, 5, 9, 9, 2}
	t.Log("原始数组")
	sisyphus.PrintSlice(data, " ")
	var err error
	data, err = sisyphus.RemoveSlice(data, 11)
	if err != nil {
		fmt.Println("Error:", err)
	}
	t.Log("调用 RemoveSlice 之后")
	sisyphus.PrintSlice(data, " ")
}

func TestDeduplicateSlice(t *testing.T) {
	data := []int{1, 4, 1, 2, 2, 5, 3, 4, 5, 9, 9, 2}
	sisyphus.PrintSlice(data, " ")

	data1 := sisyphus.DeduplicateSlice(data)
	sisyphus.PrintSlice(data1, " ")
}

func TestFibonacciSlice(t *testing.T) {
	data := sisyphus.FibonacciSlice(10)
	sisyphus.PrintSlice(data, ", ")
}

func TestFibonacci(t *testing.T) {
	data := sisyphus.Fibonacci(10)
	fmt.Println(data)
}

func TestMergeSort(t *testing.T) {
	data := []int{6, 5, 12, 10, 9, 1}
	sisyphus.PrintSlice(data, ", ")
	data = sisyphus.MergeSort(data)
	sisyphus.PrintSlice(data, ", ")
}
