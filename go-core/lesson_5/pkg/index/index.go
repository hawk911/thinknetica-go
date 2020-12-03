package index

import (
	"fmt"
	"strings"
)

type BTree struct {
	root *Element
}
type Element struct {
	left  *Element
	right *Element
	Value *Record
}
type Record struct {
	ID    int
	URL   string
	Title string
}
type Index struct {
	storage    BTree
	invIndex   map[string][]int
	IdProvider int
}

func New() *Index {
	i := Index{
		storage:    BTree{},
		invIndex:   map[string][]int{},
		IdProvider: 1,
	}
	return &i
}

func (bt *BTree) Insert(r *Record) {
	e := &Element{Value: r}
	if bt.root == nil {
		bt.root = e
		return
	}
	insert(bt.root, e)
}

func insert(node, new *Element) {
	if new.Value.ID < node.Value.ID {
		if node.left == nil {
			node.left = new
			return
		}
		insert(node.left, new)
	}
	if new.Value.ID >= node.Value.ID {
		if node.right == nil {
			node.right = new
			return
		}
		insert(node.right, new)
	}
}

// Search выполняет поиск по слову в хранилище с использованием инвертированного индекса
func (i *Index) Search(word string) []Record {
	found := []Record{}
	ids, ok := i.invIndex[word]
	if !ok {
		return found
	}

	for _, id := range ids {
		if rec, ok := i.storage.Search(id); ok {
			//found = append(found, i.storage[index])
			found = append(found, *rec)
		}
	}

	return found
}

func (i *Index) Fill(data *map[string]string) {
	for link, title := range *data {
		rec := Record{
			ID:    i.IdProvider,
			URL:   link,
			Title: title,
		}
		i.storage.Insert(&rec)

		lexemes := map[string]bool{} // map чтобы избежать дублей
		NormalizeWord(&lexemes, rec.URL)
		NormalizeWord(&lexemes, rec.Title)

		for lex := range lexemes {
			i.invIndex[lex] = append(i.invIndex[lex], rec.ID)
		}

		i.IdProvider++

	}
}

func NormalizeWord(lexemes *map[string]bool, words string) {

	words = strings.ReplaceAll(words, "://", " ")
	words = strings.ReplaceAll(words, "/", " ")
	for _, word := range strings.Fields(words) {
		word = strings.ToLower(word)
		word = strings.TrimSpace(word)
		word = strings.TrimFunc(word, func(r rune) bool {
			return (r >= 0 && r <= 64) || (r >= 91 && r <= 96)
		})

		if len([]rune(word)) > 1 {
			(*lexemes)[word] = true
		}
	}
}

// Search осуществляет поиск в бинарном дереве
func (bt *BTree) Search(id int) (*Record, bool) {
	currentNode := bt.root

	for {
		if currentNode == nil {
			return &Record{}, false
		}

		if currentNode.Value.ID == id {
			return currentNode.Value, true
		}

		if currentNode.Value.ID > id {
			currentNode = currentNode.left
			continue
		}

		currentNode = currentNode.right
	}
}

// String позволяет получить простое строковое представление бинарного дерева
func (bt *BTree) String() string {
	elems := []int{}
	bt.root.collect(&elems)
	return fmt.Sprint(elems)
}

// collect выполняет рекурсивный обход дерева и собирает Id элементов в массив
func (n *Element) collect(s *[]int) {
	if n == nil {
		return
	}

	*s = append(*s, n.Value.ID)

	n.left.collect(s)
	n.right.collect(s)
}
