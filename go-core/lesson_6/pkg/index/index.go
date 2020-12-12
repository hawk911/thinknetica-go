package index

import (
	"pkg/model"
	"strings"
)

// Index создает записи типа model.Record, формирует обратный индекс и сохраняет результаты в указанное хранилище
type Index struct {
	storage    writer
	IdProvider int
}

// writer представляет собой интерфейс (контракт) которому должно соответствовать хранилище
type writer interface {
	Write([]model.Record, model.InvIndex) error
}

// New создает новый экземпляр типа Index
func New(storage writer) *Index {
	ind := Index{
		storage:    storage,
		IdProvider: 1,
	}
	return &ind
}

//func (bt *BTree) Insert(r *Record) {
//	e := &Element{Value: r}
//	if bt.root == nil {
//		bt.root = e
//		return
//	}
//	insert(bt.root, e)
//}
//
//func insert(node, new *Element) {
//	if new.Value.ID < node.Value.ID {
//		if node.left == nil {
//			node.left = new
//			return
//		}
//		insert(node.left, new)
//	}
//	if new.Value.ID >= node.Value.ID {
//		if node.right == nil {
//			node.right = new
//			return
//		}
//		insert(node.right, new)
//	}
//}
//
//// Search выполняет поиск по слову в хранилище с использованием инвертированного индекса
//func (i *Index) Search(word string) []Record {
//	found := []Record{}
//	ids, ok := i.invIndex[word]
//	if !ok {
//		return found
//	}
//
//	for _, id := range ids {
//		if rec, ok := i.storage.Search(id); ok {
//			//found = append(found, i.storage[index])
//			found = append(found, *rec)
//		}
//	}
//
//	return found
//}

func (i *Index) Fill(data *map[string]string) error {
	records := []model.Record{}
	invIndex := model.InvIndex{}

	for link, title := range *data {
		rec := model.Record{
			ID:    i.IdProvider,
			URL:   link,
			Title: title,
		}
		records = append(records, rec)

		lexemes := map[string]bool{} // map чтобы избежать дублей
		NormalizeWord(&lexemes, rec.URL)
		NormalizeWord(&lexemes, rec.Title)

		for lex := range lexemes {
			invIndex[lex] = append(invIndex[lex], rec.ID)
		}

		i.IdProvider++

	}

	err := i.storage.Write(records, invIndex)
	if err != nil {
		return err
	}

	return nil
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

//// Search осуществляет поиск в бинарном дереве
//func (bt *BTree) Search(id int) (*Record, bool) {
//	currentNode := bt.root
//
//	for {
//		if currentNode == nil {
//			return &Record{}, false
//		}
//
//		if currentNode.Value.ID == id {
//			return currentNode.Value, true
//		}
//
//		if currentNode.Value.ID > id {
//			currentNode = currentNode.left
//			continue
//		}
//
//		currentNode = currentNode.right
//	}
//}
//
//// String позволяет получить простое строковое представление бинарного дерева
//func (bt *BTree) String() string {
//	elems := []int{}
//	bt.root.collect(&elems)
//	return fmt.Sprint(elems)
//}
//
//// collect выполняет рекурсивный обход дерева и собирает Id элементов в массив
//func (n *Element) collect(s *[]int) {
//	if n == nil {
//		return
//	}
//
//	*s = append(*s, n.Value.ID)
//
//	n.left.collect(s)
//	n.right.collect(s)
//}
