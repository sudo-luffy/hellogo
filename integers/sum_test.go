package integers

import (
	"slices"
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

func ArraySum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %d, got %d", got, want)
	}
}

func TestSum(t *testing.T) {
	sum := Sum(2, 3)
	want := 5
	if sum != want {
		t.Errorf("expected %d, got %d", want, sum)
	}
}

func TestArraySliceSum(t *testing.T) {
	t.Run("collection of 5 numbers.", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := ArraySum(numbers)
		want := 15
		assertEqual(t, got, want)
	})
	t.Run("collection of any size.", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := ArraySum(numbers)
		want := 6
		assertEqual(t, got, want)
	})
}

func SumAll(slice1 []int, slice2 []int) []int {
	sum := 0
	var all []int
	for _, num := range slice1 {
		sum += num
	}
	all = append(all, sum)
	sum = 0
	for _, num := range slice2 {
		sum += num
	}
	all = append(all, sum)
	return all
}

func SumAllVaradic(slices ...[]int) []int {
	// sums := make([]int, len(slices))
	var sums []int
	for _, slice := range slices {
		// sums[i] = ArraySum(slice)
		sums = append(sums, ArraySum(slice))
	}
	return sums
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3, 4}
	// got := SumAll(slice1, slice2)
	got := SumAllVaradic(slice1, slice2)
	want := []int{6, 10}
	if !slices.Equal(got, want) {
		t.Errorf("Expected %d, got %d", got, want)
	}
}
