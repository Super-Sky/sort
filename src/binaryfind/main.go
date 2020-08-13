package main

import "fmt"

func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int) {
	if leftIndex > rightIndex {
		fmt.Printf("找不到")
		return
	}

	middle := (leftIndex + rightIndex) /2

	if (*arr)[middle] > findVal {
		BinaryFind(arr,leftIndex,middle,findVal)
	}else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle+1,rightIndex,findVal)
	}else {
		fmt.Printf("找到下标:%v",middle)
	}
}

func main() {
	arr := [6]int{1,4,12,45,687,976}
	BinaryFind(&arr,0,len(arr)-1,45)
}
