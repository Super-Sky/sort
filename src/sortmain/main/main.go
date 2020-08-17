package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sortmain/adjustheapSort"
	"time"
)

const (
	num = 10000        //测试数组长度
	rangeNum = 100000  //数组元素大小范围
)

func GenerateRand() []int {
	randSeed := rand.New(rand.NewSource(time.Now().Unix()+time.Now().UnixNano()))
	arr := make([]int,num)
	for i :=0;i<num;i++{
		arr[i] = randSeed.Intn(rangeNum)
	}
	return arr
}

func IsSame(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	if (slice1 == nil) != (slice2 ==nil) {
		return false
	}
	for i, num := range slice1{
		if num != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	arr := GenerateRand()
	org_arr := make([]int,num)
	copy(org_arr,arr)
	//冒泡排序
	//bubbleSort.BubbleSort(arr)
	//选择排序
	//selectSort.SelectSort(arr)
	//插入排序
	//insertSort.InsertSort(arr)
	//快排
	//quickSort.QuickSort(arr,0,len(arr)-1)
	//归并排序
	//mergeSort.MergeSort(arr,0,len(arr)-1)
	//堆排序
	adjustheapSort.HeapSort(arr)
	sort.Ints(org_arr)
	fmt.Println(arr[:15],org_arr[:15],IsSame(arr,org_arr))
}

