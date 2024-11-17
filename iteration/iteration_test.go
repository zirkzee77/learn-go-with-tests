package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
  t.Run("Should print repeated char", func(t *testing.T) {
    got := Repeat("a", 5)
    want := "aaaaa"
    if (got != want) {
      t.Errorf("got %q want %q", got, want)
    }
  })
}

func ExampleRepeat() {
  result := Repeat("a", 5)
  fmt.Println(result)
  // Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
  for i:=0; i<b.N; i++ {
    Repeat("a", 5)
  }
}
