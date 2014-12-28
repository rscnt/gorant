package heapsort

func HeapSort(numbers []int) []int {
	l := len(numbers)
	start, end := (l/2)-1, l-1
	for start >= 0 {
		sift(numbers[0:], start, l)
		start -= 1
	}
	for end > 0 {
		swap(numbers[0:], end, 0)
		sift(numbers[0:], 0, end)
		end -= 1
	}
	return numbers
}

func sift(arr []int, start, count int) {
	root := start
	child := 0
	for ((root * 2) + 1) < count {
		child = ((root * 2) + 1)
		if child < count-1 && arr[child] < arr[child+1] {
			child += 1
		}
		if arr[root] < arr[child] {
			swap(arr[0:], root, child)
			root = child
		} else {
			return
		}

	}
}

func swap(arr []int, x, y int) {
	tmp := arr[y]
	arr[y] = arr[x]
	arr[x] = tmp
}
