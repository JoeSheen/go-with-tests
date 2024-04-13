package main

import "testing"

/*
test are just like functions with the following rules:
	1)	It needs to be in a file with a name like 'xxx_test.go'.
	2)	The test function must start with the word 'Test'.
	3) 	The test function takes one argument only t: '*testing.T'.
	4) 	In order to use the *testing.T type, you need to import
		"testing", like we did with fmt in the other file

*/
/*// part 2.
func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want  {
		t.Errorf("got %q want %q", got, want)
	}
}*/

/*// part 3.
func TestHello(t *testing.T) {
	got := Hello("Joe")
	want := "Hello, Joe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}*/

func TestHello(t *testing.T) {
	// The 'Run' method of T allows subtest to be defined.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joe", "")
		want := "Hello, Joe"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	// part 5.
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("Jose", "Spanish")
		want := "Hola, Jose"

		assertCorrectMessage(t, got, want)
	})

	// part 6.
	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Jacques", "French")
		want := "Bonjour, Jacques"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Swedish", func(t *testing.T) {
		got := Hello("Viktor", "Sweden")
		want := "Hallå, Viktor"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tells the test suite that this is a helper method
	
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}