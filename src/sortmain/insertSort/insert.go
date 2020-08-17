package insertSort

func InsertSort(arr []int)  {
	for i := 1;i<len(arr)-1;i++{
		for j:=i;j>0;j-- {
			if arr[j-1] >arr[j] {
				arr[j-1],arr[j] = arr[j],arr[j-1]
			}
		}
	}
}