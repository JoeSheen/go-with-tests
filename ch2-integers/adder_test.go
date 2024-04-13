package ch2integers

import (
	"fmt"
	"testing"
)

// (x int, y int) can be shorten to (x, y int) when arguments are of the same type.
func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}