package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("", func(t *testing.T) {
		repeated := Repeat("a", 15)
		expected := "aaaaaaaaaaaaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 15)
	}
}
func ExampleRepeat() {
	result := Repeat("a", 4)
	fmt.Println(result)
	//OUTPUT: aaaa
}
