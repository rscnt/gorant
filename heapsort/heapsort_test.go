package heapsort

import "testing"

var oArr = []int{1, 2, 3}
var uArr = []int{2, 1, 3}

func TestHeapSort(t *testing.T) {
	rArr := HeapSort(uArr)
	if slicesAreTheSame(rArr, oArr) {
		t.Error("Got: ", rArr,
			"Expected: ", oArr)
	}
}

func slicesAreTheSame(a, b []int) bool {
	result := false
	if len(a) == len(b) {
		for i, v := range a {
			if v != b[i] {
				result = true
			}
		}
	}
	return result
}
