package mergeSort

//合并
func Merge(arr []int, l, mid, r int) {
	// 分别复制左右子数组
	n1, n2 := mid-l+1, r-mid
	left, right := make([]int, n1), make([]int, n2)
	copy(left, arr[l:mid+1])
	copy(right, arr[mid+1:r+1])
	i, j := 0, 0
	k := l
	for; i < n1 && j < n2; k++ {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
	}
	for; i < n1; i++ {
		arr[k] = left[i]
		k++
	}
	for; j < n2; j++ {
		arr[k] = right[j]
		k++
	}
}

//分治
func MergeSort(arr []int, l, r int) {
	if l < r {
		mid := (l + r - 1) / 2
		MergeSort(arr, l, mid)
		MergeSort(arr, mid+1, r)
		Merge(arr, l, mid, r)
	}
}
