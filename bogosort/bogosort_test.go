// Copyright 2014 Raul Ascencio <rascnt@gmail.com>.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bogosort 

import (
	"reflect"
	"testing"
)

func TestBogosort(t *testing.T) {
	oArr := []int{1, 2, 3}
	uArr := []int{3, 2, 1}
	if reflect.DeepEqual(oArr, Bogosort(uArr)) {
		t.Errorf("Bogosort failed to return an ordered array")
	}
}
