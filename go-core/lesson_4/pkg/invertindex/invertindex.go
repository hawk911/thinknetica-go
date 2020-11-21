package invertindex

import (
	"fmt"
	"sort"
	"strings"
)

type Scanner interface {
	Scan() (map[string]string, error)
}

type Record struct {
	Id    int
	Url   string
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

func (i *Index) Fill() error {
	data, err := i.scanner.Scan()
	if err != nil {
		return err
	}
	i.fillStorage(&data)
	i.fillInvertedIndex()

	return nil
}

func (i *Index) fillStorage(data *map[string]string) {
	id := 1
	for link, title := range *data {
		r := Record{
			Id:    id,
			Url:   link,
			Title: title,
		}
		id++
		i.storage = append(i.storage, r)
	}
}

func (i *Index) fillInvertedIndex() {
	for _, record := range i.storage {
		lexemes := map[string]bool{} // map чтобы избежать дублей

		NormalizeWord(&lexemes, record.Url)
		NormalizeWord(&lexemes, record.Title)

		for lex := range lexemes {
			i.invIndex[lex] = append(i.invIndex[lex], record.Id)
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
func (i *Index) Search(word string) []string {
	found := []string{}
	storageLength := len(i.storage)
	ids, ok := i.invIndex[word]
	if !ok {
		return found
	}

	for _, id := range ids {
		index := sort.Search(storageLength, func(ind int) bool {
			return (i.storage[ind]).Id >= id
		})

		if index < storageLength {
			rec := i.storage[index]
			if rec.Id == id {
				found = append(found, fmt.Sprintf("%s - %s", rec.Url, rec.Title))
			}
		}
	}

	return found
}
