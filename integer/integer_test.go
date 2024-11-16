package integer

import (
	"fmt"
	"testing"
)

func TestAdd (t *testing.T) {
  got := Add(2, 2)
  want := 4
  assertExpectation(t, got, want)
}

func ExampleAdd() {
  sum := Add(2, 3)
  fmt.Println(sum)
  // Output: 5
}

func assertExpectation (t testing.TB, got int, want int) {
  t.Helper()
  if (got != want) {
    t.Errorf("got %d want %d", got, want)
  }
}
