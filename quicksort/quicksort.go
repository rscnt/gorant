package quicksort

func swap(slice []int, a, b int) {
	tmp := slice[a]
	slice[a] = slice[b]
	slice[b] = tmp
}

func diffInt(a, b int) int {
	return a - b
}

func partition(slice []int, left, right int) int {
	cmp, minEnd, maxEnd := slice[right-1], left, 0
	for maxEnd = left; maxEnd < right-1; maxEnd++ {
		if diffInt(slice[maxEnd], cmp) < 0 {
			swap(slice[0:], maxEnd, minEnd)
			minEnd++
		}
	}
	swap(slice[0:], minEnd, right-1)
	return minEnd
}

func QuickSort(slice []int, start, end int) []int {
	if start < end {
		part := partition(slice[0:], start, end)
		QuickSort(slice[0:], start, part)
		QuickSort(slice[0:], part+1, end)
	}
	return slice
}
