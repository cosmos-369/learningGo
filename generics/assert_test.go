package generics

import (
	"testing"
)

func TestAssertFunction(t *testing.T) {
	t.Run("asserts on integer", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("assert on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
	})
}

func TestStack(t *testing.T) {
	t.Run("test stack of integers", func(t *testing.T) {
		intStack := Stack[int]{}

		//check empty
		AssertTrue(t, intStack.IsEmpty())

		//check push
		intStack.Push(10)
		AssertFalse(t, intStack.IsEmpty())

		//check pop
		intStack.Push(20)
		if i, _ := intStack.Pop(); i != 20 {
			t.Errorf("did not pop recent value, got %+v want %+v", i, 20)
		}
	})

	// t.Run("test stack of strings", func(t *testing.T) {
	// 	stringStack := Stack[string]{}

	// 	//check empty
	// 	AssertTrue(t, stringStack.IsEmpty())

	// 	//check push
	// 	stringStack.Push("hello")
	// 	AssertFalse(t, stringStack.IsEmpty())
	// 	//check pop
	// 	stringStack.Push("world")
	// 	val, _ := stringStack.Pop()
	// 	if val != "world" {
	// 		t.Errorf("did not pop recent value, got %+v want %+v", val, "world")
	// 	}
	// })

	t.Run("return error when try to Pop empty Stack", func(t *testing.T) {
		newStack := Stack[int]{}

		_, err := newStack.Pop()

		if err == true {
			t.Error("expected false but got true")
		}
	})

	t.Run("interface stack DX is horrid", func(t *testing.T) {
		myStackOfInts := Stack[int]{}

		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func AssertNotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %+v, want %+v", got, want)
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
