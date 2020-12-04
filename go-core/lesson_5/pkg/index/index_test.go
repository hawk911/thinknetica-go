package index

import (
	"pkg/crawbot"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := crawbot.New()
	data, _ := scanner.Scan()
	ind := New()
	ind.Fill(&data)

	found := ind.Search("http")
	got := len(found)
	want := 3
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = ind.Search("golang.org")
	got = len(found)
	want = 1
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}

	found = ind.Search("wish")
	got = len(found)
	want = 0
	if got != want {
		t.Errorf("len(found) = %d; want %d", got, want)
	}
}
