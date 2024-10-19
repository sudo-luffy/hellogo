package iter

import "testing"

func assertEqual(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("Expected %s got %s", got, want)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat("a", 4)
	want := "aaaa"

	assertEqual(t, got, want)
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
