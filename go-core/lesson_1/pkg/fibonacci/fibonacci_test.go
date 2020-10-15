package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {
	got := Fibonacci(10)
	want := 55
	if got != want {
		t.Errorf("Fibonacci(10) = %d; want %d", got, want)
	}

	got = Fibonacci(1)
	want = 1
	if got != want {
		t.Errorf("Fibonacci(10) = %d; want %d", got, want)
	}

	got = Fibonacci(20)
	want = 6765
	if got != want {
		t.Errorf("Fibonacci(10) = %d; want %d", got, want)
	}
}
