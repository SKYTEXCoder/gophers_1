package greeting

import "testing"

func TestHello(t *testing.T) {
	got := Hello("World")
	want := "Hello, World! Welcome to go_task1."
	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
