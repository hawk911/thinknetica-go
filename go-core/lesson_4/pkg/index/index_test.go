package index

import (
	"fmt"
	"pkg/crawbot"
	"testing"
)

func TestSearch(t *testing.T) {
	scanner := crawbot.New()
	ind := New(scanner)
	data, err := ind.Fill()
	if err != nil {
		fmt.Println(err)
		return
	}
	ind.FillStorage(&data)
	ind.FillInvertedIndex()

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
