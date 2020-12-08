// Package model предоставляет структуры данных которые используются в нескольких пакетах приложения
package model

type Record struct {
	ID    int
	URL   string
	Title string
}

type InvIndex map[string][]int
