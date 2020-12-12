// Package engine читает данные из хранилища и предоставляет к ним доступ
package engine

import (
	"pkg/btree"
	"pkg/model"
)

// Engine предоставляет методы для поиска данных в указанном хранилище
type Engine struct {
	storage       storage
	tree          btree.BTree
	invertedIndex model.InvIndex
}

// storage представляет собой интерфейс (контракт) которому должно соответствовать хранилище
type storage interface {
	Read() ([]model.Record, model.InvIndex, error)
	IsUpdated() bool
}

// New создает новый экземпляр типа Engine
func New(s storage) *Engine {
	eng := Engine{
		storage: s,
		tree:    btree.BTree{},
	}
	return &eng
}

// Search осуществляет поиск по слову
func (e *Engine) Search(word string) ([]model.Record, error) {
	if e.storage.IsUpdated() {
		err := e.update()
		if err != nil {
			return []model.Record{}, err
		}
	}

	found := []model.Record{}
	ids, ok := e.invertedIndex[word]
	if !ok {
		return found, nil
	}

	for _, id := range ids {
		if rec, ok := e.tree.Search(id); ok {
			found = append(found, *rec)
		}
	}

	return found, nil
}

// update читает данные из хранилища и строит в памяти необходимые структуры для быстрого доступа к ним
func (e *Engine) update() error {
	records, index, err := e.storage.Read()
	if err != nil {
		return err
	}

	for _, rec := range records {
		record := rec // приходится копировать здесь, иначе все записи ссылаются на одно и то же значение
		e.tree.Add(&record)
	}
	e.invertedIndex = index

	return nil
}
