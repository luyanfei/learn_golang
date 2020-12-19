package handy_test

import (
	. "handy"
	"reflect"
	"testing"
)

func TestRepeatAndTake(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	expected := []int{1, 2, 1, 2, 1}
	result := make([]int, 0)
	for v := range Take(done, Repeat(done, 1, 2), 5) {
		if iv, ok := v.(int); ok {
			result = append(result, iv)
		}
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result(%v) is not equals to expected(%v).", result, expected)
	}
}
