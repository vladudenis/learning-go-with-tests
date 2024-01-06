package iteration

import (
	"strings"
	"testing"
)

func assertCorrectMessage(t testing.TB, exp, res string) {
	t.Helper()

	if res != exp {
		t.Errorf("Result: %q; Desired Result: %q", res, exp)
	}
}

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"
		assertCorrectMessage(t, expected, repeated)
	})
	t.Run("repeat 0 times", func(t *testing.T) {
		repeated := Repeat("a", 0)
		expected := ""
		assertCorrectMessage(t, expected, repeated)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 100)
	}
}

func BenchmarkStringRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 100)
	}
}
