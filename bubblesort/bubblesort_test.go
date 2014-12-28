package bubblesort

import (
	"testing"
)

var oArr = []int{1, 2, 3}
var uArr = []int{2, 3, 1}

func TestOrder(t *testing.T) {
	tmpArr := BubbleSort(uArr)
	if !areEqualIntArr(oArr, tmpArr) {
		t.Error("Got: ", tmpArr,
			"Expected: ", oArr,
		)
	}

}

func areEqualIntArr(xArr, yArr []int) bool {
	result := true
	if len(xArr) != len(yArr) {
		for i, _ := range xArr {
			if xArr[i] != yArr[i] {
				return false
			}
		}
	}
	return result
}
