package positive_test

import (
	. "positive"
	"testing"
)

func TestPositive(t *testing.T) {
	if val1, _ := Positive(1); !val1 {
		t.Errorf("Positive(1) is not true")
	}
	if val2, _ := Positive(-1); val2 {
		t.Errorf("Positive(-1) is true")
	}
	_, err := Positive(0)
	if err == nil {
		t.Errorf("Positive(0) has no error")
	}
	if err.Error() != "undefined" {
		t.Errorf("Positive(0) error do not has message 'undefined'")
	}
}
