package memory

import (
	"fmt"
	"pkg/model"
	"testing"
)

func TestWrite_Read(t *testing.T) {
	store, err := NewStorage()
	if err != nil {
		t.Errorf("NewStorage(); err = %s; want nil", err)
		return
	}

	records := []model.Record{
		model.Record{
			ID:    1,
			URL:   "http://one.com",
			Title: "First title",
		},
		model.Record{
			ID:    2,
			URL:   "http://two.com",
			Title: "Second title",
		},
	}
	index := model.InvIndex{}
	index["one"] = []int{1, 3}
	index["two"] = []int{1, 2}

	err = store.Write(records, index)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}
	readRecords, readIndex, err := store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want := records[0]
	got := readRecords[0]
	if got != want {
		t.Errorf("readRecords[0] = %v; want %v", got, want)
	}
	want = records[1]
	got = readRecords[1]
	if got != want {
		t.Errorf("readRecords[1] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index["one"])
		got := fmt.Sprint(readIndex["one"])
		if got != want {
			t.Errorf("readIndex[\"one\"] = %s; want %s", got, want)
		}
		want = fmt.Sprint(index["two"])
		got = fmt.Sprint(readIndex["two"])
		if got != want {
			t.Errorf("readIndex[\"two\"] = %s; want %s", got, want)
		}
	}

	records = []model.Record{
		model.Record{
			ID:    3,
			URL:   "http://three.com",
			Title: "Third title",
		},
	}
	index = model.InvIndex{}
	index["three"] = []int{4, 5}

	err = store.Write(records, index)
	if err != nil {
		t.Errorf("store.Write(); err = %s; want nil", err)
		return
	}

	readRecords, readIndex, err = store.Read()
	if err != nil {
		t.Errorf("store.Read(); err = %s; want nil", err)
		return
	}
	want = records[0]
	got = readRecords[0]
	if got != want {
		t.Errorf("readRecords[0] = %v; want %v", got, want)
	}
	{
		want := fmt.Sprint(index["three"])
		got := fmt.Sprint(readIndex["three"])
		if got != want {
			t.Errorf("readIndex[\"three\"] = %s; want %s", got, want)
		}
	}
}
