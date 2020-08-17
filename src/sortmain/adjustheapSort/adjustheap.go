package adjustheapSort

//堆调整
func adjust_heap(arr []int, i, size int) {
	if i <= (size-2)/2 {
		//左右子节点
		l, r := 2*i+1, 2*i+2
		m := i
		if l < size && arr[l] > arr[m] {
			m = l
		}
		if r < size && arr[r] > arr[m] {
			m = r
		}
		if m != i {
			arr[m], arr[i] = arr[i], arr[m]
			adjust_heap(arr, m, size)
		}
	}
}

//建堆
func build_heap(arr []int) {
	size := len(arr)
	//从最后一个子节点开始向前调整
	for i := (size - 2) / 2; i >= 0; i-- {
		adjust_heap(arr, i, size)
	}
}

func HeapSort(arr []int) {
	size := len(arr)
	build_heap(arr)
	for i := size - 1; i > 0; i-- {
		//顶部arr[0]为当前最大值,调整到数组末尾
		arr[0], arr[i] = arr[i], arr[0]
		adjust_heap(arr, 0, i)
	}
}
