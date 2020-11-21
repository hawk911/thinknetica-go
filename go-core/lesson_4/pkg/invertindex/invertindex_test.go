package invertindex

import (
	"pkg/crawbot"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := crawbot.New()
	i := New(scanner)
	i.Fill()

	found := i.Search("http")
	got := len(found)
	want := 3
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = i.Search("golang.org")
	got = len(found)
	want = 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = i.Search("wish")
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
