package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
		// AssertEqual(t, 1, "1") // this is invalid and will error
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Joe")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		// check stack is empty when first created
		AssertTrue(t, stackOfInts.IsEmpty())

		// add an int, then check it's not empty
		stackOfInts.Push(123)
		AssertFalse(t, stackOfInts.IsEmpty())

		// add another int, pop it back again
		stackOfInts.Push(456)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, stackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		stackOfInts.Push(1)
		stackOfInts.Push(2)
		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %v", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
