package selectSort

func SelectSort(arr []int)  {
	for i :=0;i<len(arr)-1;i++ {
		for j :=i+1; j<=len(arr)-1;j++{
			if (arr)[j] < (arr)[i] {
				(arr)[j],(arr)[i] = (arr)[i],(arr)[j]
			}
		}
	}
}
