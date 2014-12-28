// Copyright 2014 Raul Ascencio <rascnt@gmail.com>.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bubblesort

func BubbleSort(arr []int) []int {
	sorted := false
	for !sorted {
		sorted = true
		for i := len(arr) - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}
