package engine

import (
	"pkg/crawbot"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := crawbot.New()

	found, err := Search(scanner, "There")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got := len(found)
	want := 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found, err = Search(scanner, "world")
	if err != nil {
		t.Errorf("err = %s; want nil", err)
	}
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
