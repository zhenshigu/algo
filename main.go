package main

import (
	"fmt"
	"math/rand"
	"sort"

)

func main() {
	a := []int{2, 4, 1, 3, 5, 7, 31, 15, 6, 33, 21, 1, 8, 98, 19, 100}
	index:=sort.Search(0, func(i int) bool {
		return a[i]>6
	})
	fmt.Println(index)
}

//10,基数排序
func radixSort(arr []int) []int {
	max := arr[0]
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	for exp := 1; max/exp > 0; exp *= 10 {
		counterArray := make([]int, 10)
		for _, num := range arr {
			counterArray[(num/exp)%10]++
		}
		for i := 1; i < 10; i++ {
			counterArray[i] += counterArray[i-1]
		}
		sortedArray := make([]int, len(arr))
		for j := len(arr) - 1; j >= 0; j-- {
			sortedArray[counterArray[(arr[j]/exp)%10]-1] = arr[j]
			counterArray[(arr[j]/exp)%10]--
		}
		arr = sortedArray
		counterArray = nil
		sortedArray = nil
	}
	return arr
}

//9,桶排序
func bucketSort(arr []int) []int {
	bucketSize := 5
	min := arr[0]
	max := arr[0]
	for _, item := range arr {
		if item < min {
			min = item
		}
		if item > max {
			max = item
		}
	}
	bucketCount := (max-min)/bucketSize + 1
	buckets := make([][]int, bucketCount)
	for _, item := range arr {
		buckets[item/bucketSize] = append(buckets[item/bucketSize], item)
	}
	result := make([]int, len(arr))
	for _, bucket := range buckets {
		bucket = selectSort(bucket)
		result = append(result, bucket...)
	}
	return result
}

//8,计数排序
func countSort(arr []int, max int) []int {
	buckets := make([]int, max+1)
	for _, item := range arr {
		buckets[item]++
	}
	sorted := make([]int, len(arr))
	j := 0
	for i := 0; i <= max; i++ {
		for buckets[i] > 0 {
			sorted[j] = i
			j++
			buckets[i]--
		}
	}
	return sorted
}

//7,快速排序
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	lowPart := make([]int, 0, len(arr))
	samePart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	pivot := arr[rand.Intn(len(arr))]
	for _, item := range arr {
		if item < pivot {
			lowPart = append(lowPart, item)
		} else if item > pivot {
			highPart = append(highPart, item)
		} else {
			samePart = append(samePart, item)
		}
	}
	lowPart = quickSort(lowPart)
	lowPart = append(lowPart, samePart...)
	highPart = quickSort(highPart)
	lowPart = append(lowPart, highPart...)
	return lowPart
}

//6,堆排序
func heapSort(arr []int) []int {
	buildHeap(arr)
	n := len(arr) - 1
	for n > 1 {
		arr[1], arr[n] = arr[n], arr[1]
		n--
		heapify(arr, 1, n)
		fmt.Println(arr)
	}
	return arr
}

//堆化
func buildHeap(arr []int) {
	n := len(arr) - 1
	for i := n / 2; i > 0; i-- {
		heapify(arr, i, n)
	}
}
func heapify(arr []int, index int, n int) {
	for {
		maxIndex := index
		if 2*index <= n && arr[index] < arr[2*index] {
			maxIndex = 2 * index
		}
		if 2*index+1 <= n && arr[maxIndex] < arr[2*index+1] {
			maxIndex = 2*index + 1
		}
		if maxIndex == index {
			fmt.Println(arr)
			break
		}
		arr[maxIndex], arr[index] = arr[index], arr[maxIndex]
		index = maxIndex
	}

}

//5,选择排序
func selectSort(arr []int) []int {
	for i := 0; i < (len(arr) - 1); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

//4,冒泡排序
func bubbleSort(arr []int) []int {
	s := len(arr) - 1
	for i := 0; i < (len(arr) - 1); i++ {
		for j := 0; j < s; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		s--
	}
	return arr
}

//3,插入排序
func insertSort(arr []int) []int {
	if arr == nil {
		return arr
	}
	b := make([]int, 0, 10)
	b[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		j := i - 1
		for j > 0 && b[j] > arr[i] {
			b[j+1] = b[j]
			j--
		}
		b[j+1] = arr[i]
	}
	return b
}

//2,归并排序
func mergeSort(arr []int, start int, end int) []int {
	if start == end {
		return arr[start : start+1]
	}
	mid := (start + end) / 2
	sortedA := mergeSort(arr, start, mid)
	sortedB := mergeSort(arr, mid+1, end)
	return sort1(sortedA, sortedB)
}
func sort1(a []int, b []int) []int {
	result := make([]int, 0, 10)
	if a[len(a)-1] < b[0] {
		result = append(result, a...)
		result = append(result, b...)
	} else if b[len(b)-1] < a[0] {
		result = append(result, b...)
		result = append(result, a...)
	} else {
		i := 0
		j := 0
		for i < len(a) && j < len(b) {
			if a[i] < b[j] {
				result = append(result, a[i])
				i++
			} else {
				result = append(result, b[j])
				j++
			}
		}
		result = append(result, a[i:]...)
		result = append(result, b[j:]...)
	}
	return result

}

//1，希尔排序
func shellSort(arr []int) []int {
	l := len(arr)
	for gap := l / 2; gap > 0; gap /= 2 {
		for i := gap; i < l; i++ {
			shellInsert(arr, gap, i)
		}
	}
	return arr
}
func shellInsert(arr []int, gap int, i int) {
	var last int
	inserted := arr[i]
	for last = i - gap; last >= 0 && inserted < arr[last]; last -= gap {
		arr[last+gap] = arr[last]
	}
	arr[last+gap] = inserted
}
