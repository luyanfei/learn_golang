package pipeline_test

import (
	"fmt"
	. "pipeline"
	"reflect"
	"testing"
)

func TestPipeline(t *testing.T) {
	arr := make([]int, 0)
	done := make(chan interface{})
	defer close(done)
	intStream := Generator(done, 1, 2, 3, 4)
	p := Multiply(done, Add(done, Multiply(done, intStream, 2), 1), 2)
	for v := range p {
		fmt.Println(v)
		arr = append(arr, v)
	}
	expected := []int{6, 10, 14, 18}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("result(%v) is not expected(%v).", arr, expected)
	}
}
