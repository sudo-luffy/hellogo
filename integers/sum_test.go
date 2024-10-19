package integers

import (
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	sum := Sum(2, 3)
	want := 5
	if sum != want {
		t.Errorf("expected %d, got %d", want, sum)
	}
}
