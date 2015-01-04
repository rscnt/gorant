package quicksort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var uArr = []int{2, 1, 3}
var oArr = []int{1, 2, 3}

func TestQuickSort(t *testing.T) {
	result := QuickSort(uArr[0:], 0, len(uArr))
	assert.Equal(t, result, oArr, "Result should be equal to oArr")
}
