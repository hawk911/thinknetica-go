package index

import (
	"sort"
	"strings"
)

type Scanner interface {
	Scan() (map[string]string, error)
}

var id int

type Record struct {
	ID    int
	URL   string
	Title string
}

type Index struct {
	scanner  Scanner
	storage  []Record
	invIndex map[string][]int
}

func New(s Scanner) *Index {
	i := Index{
		scanner:  s,
		storage:  []Record{},
		invIndex: map[string][]int{},
	}
	return &i
}

func (i *Index) Fill() (map[string]string, error) {
	data, err := i.scanner.Scan()
	if err != nil {
		return data, err
	}
	return data, err
}

func (i *Index) FillStorage(data *map[string]string) {
	for link, title := range *data {
		r := Record{
			ID:    id,
			URL:   link,
			Title: title,
		}
		id++
		i.storage = append(i.storage, r)
	}
}

func (i *Index) FillInvertedIndex() {
	for _, record := range i.storage {
		lexemes := map[string]bool{} // map чтобы избежать дублей

		NormalizeWord(&lexemes, record.URL)
		NormalizeWord(&lexemes, record.Title)

		for lex := range lexemes {
			i.invIndex[lex] = append(i.invIndex[lex], record.ID)
		}
	}
}

func NormalizeWord(lexemes *map[string]bool, words string) {

	words = strings.ReplaceAll(words, "://", " ")
	words = strings.ReplaceAll(words, "/", " ")
	for _, word := range strings.Fields(words) {
		word = strings.ToLower(word)
		word = strings.TrimSpace(word)
		word = strings.TrimFunc(word, func(r rune) bool {
			return ((r >= 0 && r <= 64) || (r >= 91 && r <= 96) || (r >= 123))
		})

		if len([]rune(word)) > 1 {
			(*lexemes)[word] = true
		}
	}
}

// Search выполняет поиск по слову в хранилище с использованием инвертированного индекса
// используется sort.Search(), который требует дополнительной проверки что коллекция содержит искомый элемент
func (i *Index) Search(word string) []Record {
	found := []Record{}
	storageLength := len(i.storage)
	ids, ok := i.invIndex[word]
	if !ok {
		return found
	}

	for _, id := range ids {
		index := sort.Search(storageLength, func(ind int) bool {
			return (i.storage[ind]).ID >= id
		})

		if index < storageLength {
			rec := i.storage[index]
			if rec.ID == id {
				found = append(found, i.storage[index])
			}
		}
	}

	return found
}
