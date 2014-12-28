// Copyright 2014 Raul Ascencio <rascnt@gmail.com>.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bogosort

import (
	"math/rand"
)

// Bogosort sorts an integer array in ascending order, the array passed
// must not be a null reference.

func shuffle(src []int) []int {
	dest := make([]int, len(src))
	perm := rand.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}

func isSorted(src []int) bool {
	result := true
	for i := 1; i < len(src); i++ {
		if src[i] < src[i-1] {
			result = false
		}
	}
	return result
}

func Bogosort(arr []int) []int {
	for !isSorted(arr) {
		arr = shuffle(arr)
	}
	return arr
}
