package bubbleSort

//冒泡排序：
//通过对序列从后往前,依次比较相邻元素的排序码，若发现逆序则交换，使排序码较小的元素逐渐从后向前移动。


func BubbleSort(arr []int) {
	lenArr := len(arr)
	for i:=0;i<lenArr-1;i++ {
		for j:=0;j<lenArr-i-1;j++ {
			if (arr)[j] > (arr)[j+1] {
				(arr)[j],(arr)[j+1] = (arr)[j+1],(arr)[j]
			}
		}
	}
}

//func main() {
//	arr := [5]int{5,1,9,24,11}
//	fmt.Printf("排序前：%v\n",arr)
//	bubbleSort(&arr)
//	fmt.Printf("排序后：%v\n",arr)
//}